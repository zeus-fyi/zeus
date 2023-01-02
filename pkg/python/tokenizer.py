import sys
import transformers

# Load the GPT-2 tokenizer
tokenizer = transformers.OpenAIGPTTokenizer.from_pretrained('openai-gpt')

# Encode a string using the GPT-2 tokenizer
prompt = sys.argv[1]
encoded_input = tokenizer.encode(prompt, return_tensors='pt')

# Print the encoded input
count = 0
for tokenlist in encoded_input:
    count += len(tokenlist)

print(count)
