import gradio as gr
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer, StoppingCriteria, StoppingCriteriaList, \
    TextIteratorStreamer
from transformers.generation.utils import GenerationConfig
from threading import Thread
from peft import PeftModel
from registory import my_chat_interface as chat_interface
import registory.utils

model_path = "/LLM/baichuan/Baichuan2-7B-Chat"
lora_path = "/LLM/baichuan/FT/A100-7B-Chat-2048-64-16-16-2-1e-5-constant_self_instruct_eval-a/checkpoint-797/"

print("init model...")
model = AutoModelForCausalLM.from_pretrained(
    model_path,
    torch_dtype=torch.float16,
    device_map="cuda:1",
    trust_remote_code=True
)
model = PeftModel.from_pretrained(model, lora_path)
model.generation_config = GenerationConfig.from_pretrained(model_path)
tokenizer = AutoTokenizer.from_pretrained(
    model_path,
    use_fast=False,
    trust_remote_code=True
)


def generate_response(message, history):
    try:
        history_formatted = history + [[message, ""]]
        messages = "\n".join(["<human>:" + item[0] + "\n<bot>:" + item[1] for item in history_formatted])
        print(messages)

        inputs = tokenizer.encode(messages, return_tensors="pt").half().to("cuda")

        generate_kwargs = {
            "max_length": 2048 + inputs.shape[1],
            "temperature": 0.7,
            "top_p": 0.95,
            "top_k": 50,
            "num_return_sequences": 1,
            "pad_token_id": tokenizer.eos_token_id,
        }

        streamer = TextIteratorStreamer(tokenizer, timeout=10., skip_prompt=True, skip_special_tokens=True)

        t = Thread(target=model.generate, kwargs=generate_kwargs).start()

        # generate_sequence = outputs[0].cpu().tolist()
        # bot_message = tokenizer.decode(generate_sequence, clean_up_tokenization_spaces=True)

        # return bot_message.replace("<|endoftext|>", "")

        partial_message = ""
        for new_token in streamer:
            if new_token != tokenizer.eos_token:
                partial_message += new_token
                yield partial_message

    except Exception as e:
        print(f"Error during generation: {e}")
        return "I'm sorry, I cannot generate a response at the moment."


def predict(message, history, request: gr.Request):
    generate_response(message, history)
    # registory.utils.record_question(request, message, bot_message)


if __name__ == "__main__":
    chat_interface.ChatInterface(
        fn=predict,
        chatbot=gr.Chatbot(scale=1, height=200),
        css="margin.css",
        undo_btn=None,
    ).launch(server_name="10.30.19.40", server_port=8082)
