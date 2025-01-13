# TermAI

A Terminal based application written in Golang for quick Q&A with GenAI models.

# Supported Models

Currently supports Gemini (1.5 Flash) and OpenAI (GPT-4o)

# Clone and install the application
# Important
> You Need To Have [Go](https://go.dev/) Installed To Use This Application
1. Clone the repo using the following command
   ```
   git clone https://github.com/srivats22/termai.git
   ```
2. Open a terminal/command prompt window and navigate to where the project is cloned
3. Run the following command
   ```
   go install
   ```
4. Once the install is completed run the following command
   ```
   termai
   ```
   this will generate an output that looks like the image below
   <img width="584" alt="Screenshot 2025-01-13 at 7 50 50 AM" src="https://github.com/user-attachments/assets/a7f06ec5-a72b-49bd-ba87-d36c923b5c4a" />

# Initial Setup

Once you have cloned and installed the application, you can move on to the one-time setup of AI Models.

1. Run the following command
   ```termai setup```
   you will see a screen like below choose the AI model you want to configure first
   
   ![image](https://github.com/user-attachments/assets/2ebc6641-0b1d-4a0c-bd9a-155a100761e0)

   
2. After Choosing the Model you will be prompted to enter you API key once you complete that step you can start interacting with the models using the command mentioned invoking the AI Models
![image](https://github.com/user-attachments/assets/23e1cd5f-7135-40d4-b8d9-91bcff4d70f2)


# Invoking The AI Models
1. To Invoke Gemini
   > Replace Enter Question with the question you want to ask Gemini
   > This command supports streaming so you will see the response as soon as the AI starts replying
   ```
   termai gemini Enter Question
   ```
2. To Invoke OpenAI
   > Replace Enter Question with the question you want to ask OpenAI
   > Currently the application doesn't support streaming and will be made available in the future
   ```
   termai oai Enter Question
   ```
