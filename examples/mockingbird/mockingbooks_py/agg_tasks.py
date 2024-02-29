import json

import requests

from examples.mockingbird.mockingbooks_py.api_setup import api_v1_path, get_headers


def create_agg_task(task):
    url = api_v1_path + "/tasks/ai"
    headers = get_headers()
    response = requests.post(url, json=task, headers=headers)
    # Check the response status
    if response.status_code == 200:
        print("Task created successfully!")
    else:
        print(response.json())
        print("Failed to create task. Status Code:", response.status_code)
    return response.json()


if __name__ == '__main__':
    with open('templates/aggregation.json', 'r') as file:
        payload = json.load(file)

    payload['taskName'] = 'demo-agg-task'
    payload['taskGroup'] = 'demo'
    payload['model'] = 'gpt-3.5-turbo-0125'
    payload['prompt'] = ('how do you think AI will respond to being offered a'
                         ' decision tree with cost assigned options and a budget?')

    create_agg_task(payload)
