import json  # Import the json module

import requests

from examples.mockingbird.mockingbooks_py.api_setup import get_headers, api_v1_path


def get_retrieval(rid):
    headers = get_headers()
    url = api_v1_path + "/retrieval/ai/" + rid
    response = requests.get(url, headers=headers)

    if response.status_code == 200:
        data = response.json()
        with open('ret-' + rid + '.json', 'w') as json_file:
            json.dump(data, json_file, indent=4)
        pretty_data = json.dumps(data, indent=4)
        print(pretty_data)
    else:
        print("Failed to fetch task data. Status Code:", response.status_code)
    return response.json()


def get_retrievals():
    headers = get_headers()
    url = api_v1_path + "/retrievals/ai"
    response = requests.get(url, headers=headers)

    if response.status_code == 200:
        return response.json()
    else:
        print("Failed to fetch task data. Status Code:", response.status_code)
    return response.json()


def get_retrieval_by_name(ret_name):
    rets = get_retrievals()
    for ret in rets:
        if ret['retrievalName'] == ret_name:
            return ret
    return None


def get_retrieval_id_by_name(ret_name):
    ret = get_retrieval_by_name(ret_name)
    if ret:
        return ret['retrievalStrID']
    return '0'


def create_or_update_retrieval(ret):
    url = api_v1_path + "/retrievals/ai"
    headers = get_headers()
    response = requests.post(url, json=ret, headers=headers)
    # Check the response status
    if response.status_code == 200:
        print("Retrieval created successfully!")
    else:
        print(response.json())
        print("Failed to create retrieval. Status Code:", response.status_code)
    return response.json()


if __name__ == '__main__':
    with open('twitter/indexer_retrieval.json', 'r') as file:
        payload = json.load(file)

    payload['retrievalName'] = 'tweets-retrieval'
    payload['retrievalGroup'] = 'demo'
    payload['retrievalPlatform'] = 'twitter'
    payload['retrievalItemInstruction']['retrievalPlatformGroups'] = 'llm'

    create_or_update_retrieval(payload)
