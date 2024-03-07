import gradio as gr
from transformers import AutoTokenizer, AutoModelForCausalLM
import torch
from registory import my_chat_interface as chat_interface

# 初始化模型和tokenizer
DEVICE = "cuda" if torch.cuda.is_available() else "cpu"
MODEL_ID = "/LLM/bluelm/BlueLM-7B-Chat-32K"
tokenizer = AutoTokenizer.from_pretrained(MODEL_ID, trust_remote_code=True, use_fast=False)
model = AutoModelForCausalLM.from_pretrained(MODEL_ID, device_map={-1: DEVICE}, torch_dtype=torch.bfloat16,
                                             trust_remote_code=True).eval()

MAX_NEW_TOKENS = 512
REPETITION_PENALTY = 1.1


# 保留对话历史的预测函数
def predict_with_history(history, user_input):
    # 更新对话历史
    history.join(f"[|Human|]:{user_input}")
    # 构建包含整个历史的prompt
    prompt = "".join(history) + "[|AI|]:"
    inputs = tokenizer(prompt, return_tensors="pt", padding=True, truncation=True, max_length=1024)
    inputs = inputs.to(DEVICE)
    input_ids = inputs["input_ids"]

    # 生成回复
    output = model.generate(input_ids=input_ids, max_new_tokens=MAX_NEW_TOKENS,
                            repetition_penalty=REPETITION_PENALTY, pad_token_id=tokenizer.eos_token_id)
    response_text = tokenizer.decode(output[0], skip_special_tokens=True)

    # 从生成的文本中提取AI的最后一条回复
    response_text = response_text[len(prompt):]
    history.append(f"[|AI|]:{response_text}")

    # 仅返回AI的最后一条回复和更新的历史
    return response_text, "\n".join(history)


# 创建Gradio界面，使用状态特性来保留对话历史
iface = chat_interface.ChatInterface(
    fn=predict_with_history,
    chatbot=gr.Chatbot(scale=1, height=600),
    css="margin.css",
    undo_btn=None,
)

if __name__ == "__main__":
    iface.launch(server_name="10.30.19.49", server_port=8082)
