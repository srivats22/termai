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
   <img width="586" alt="Screenshot 2025-01-02 at 2 02 00 PM" src="https://github.com/user-attachments/assets/1482fe0b-e5f0-43d7-a107-c9be5011ed1d" />

# Initial Setup

Once you have cloned and installed the application, you can move on to the one-time setup of AI Models.

1. Run the following command
   ```termai setup```
   you will see a screen like below choose the AI model you want to configure first
   
   <img width="584" alt="Screenshot 2025-01-02 at 2 05 25 PM" src="https://github.com/user-attachments/assets/5b22f6b9-8782-4ec4-8f3a-30cebc9df0ce" />
   
2. After Choosing the Model you will be prompted to enter you API key once you complete that step you can start interacting with the models using the command mentioned invoking the AI Models
<img width="586" alt="Screenshot 2025-01-02 at 2 05 53 PM" src="https://github.com/user-attachments/assets/e0fd131f-6488-4ce2-b7be-daba2d45e12a" />

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
