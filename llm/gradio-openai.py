import openai
import gradio as gr

from registory import my_chat_interface as chat_interface
import registory.utils


def predict(message, history, request: gr.Request):
    # 将聊天历史格式化为OpenAI API所需的格式
    history_openai_format = []
    for human, assistant in history:
        history_openai_format.append({"role": "user", "content": human})
        history_openai_format.append({"role": "assistant", "content": assistant})
    # 添加最新的用户消息
    history_openai_format.append({"role": "user", "content": message})

    # 调用OpenAI API，生成聊天机器人的回应
    response = openai.ChatCompletion.create(
        model='gpt-3.5-turbo',  # 指定使用的模型
        messages=history_openai_format,
        temperature=1.0,  # 设置创造性参数
        stream=True  # 使用流式响应
    )

    # 逐步生成回复内容
    partial_message = ""
    for chunk in response:
        if 'choices' in chunk and len(chunk['choices'][0]['delta']) != 0:
            partial_message += chunk['choices'][0]['delta']['content']
            yield partial_message

    # print(request.headers)

    registory.utils.record_question(request, message, partial_message)


if __name__ == "__main__":
    chat_interface.ChatInterface(
        fn=predict,
        chatbot=gr.Chatbot(scale=1, height=200),
        css="margin.css",
        undo_btn=None,
        # theme="soft"
    ).launch(server_name="10.130.10.68", server_port=8082)
