import gradio as gr
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer, StoppingCriteria, StoppingCriteriaList, \
    TextIteratorStreamer
from transformers.generation.utils import GenerationConfig
from threading import Thread
from peft import PeftModel
from registory import my_chat_interface as chat_interface
import registory.utils

model_path = "/LLM/bluelm/BlueLM_7B_Chat_32K/"
# lora_path = "/LLM/baichuan/FT/A100-7B-Chat-2048-64-16-16-2-1e-5-constant_self_instruct_eval-a/checkpoint-797/"

print("init model...")
model = AutoModelForCausalLM.from_pretrained(
    model_path,
    torch_dtype=torch.bfloat16,
    device_map="cuda:0",
    trust_remote_code=True
).eval()
# model = PeftModel.from_pretrained(model, lora_path)
# model.generation_config = GenerationConfig.from_pretrained(model_path)
tokenizer = AutoTokenizer.from_pretrained(
    model_path,
    use_fast=False,
    trust_remote_code=True
)


class StopOnTokens(StoppingCriteria):
    def __call__(self, input_ids: torch.LongTensor, scores: torch.FloatTensor, **kwargs) -> bool:
        stop_ids = [29, 0]
        for stop_id in stop_ids:
            if input_ids[0][-1] == stop_id:
                return True
        return False


def predict(message, history, request: gr.Request):
    history_transformer_format = history + [[message, ""]]
    # stop = StopOnTokens()

    messages = "".join(["".join(["\n[|Human|]:" + item[0], "\n[|AI|]:" + item[1]])  # curr_system_message +
                        for item in history_transformer_format])

    model_inputs = tokenizer([messages], return_tensors="pt", padding=True).to("cuda")
    streamer = TextIteratorStreamer(tokenizer, timeout=10., skip_prompt=True, skip_special_tokens=True)
    generate_kwargs = dict(
        model_inputs,
        streamer=streamer,
        max_new_tokens=2048,
        do_sample=True,
        top_p=0.95,     # model从累计概率大于或等于p的最小集合中随机选择一个token
        top_k=2000,     # 保留概率最高的k个token
        temperature=0.95,
        repetition_penalty=1.1,
        num_beams=1,
        # stopping_criteria=StoppingCriteriaList([stop])
        pad_token_id=tokenizer.pad_token_id
    )
    t = Thread(target=model.generate, kwargs=generate_kwargs)
    t.start()

    partial_message = ""
    for new_token in streamer:
        if new_token != "[|Human|]:":
            partial_message += new_token
            yield partial_message

    # registory.utils.record_question(request, message, partial_message)


if __name__ == "__main__":
    chat_interface.ChatInterface(
        fn=predict,
        chatbot=gr.Chatbot(scale=1, height=600),
        css="margin.css",
        undo_btn=None,
    ).launch(server_name="10.30.19.40", server_port=8082)
