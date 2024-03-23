import json  # Import the json module

import requests

from examples.mockingbird.mockingbooks_py.api_setup import get_headers, api_v1_path


def get_triggers():
    url = api_v1_path + "/actions/ai"
    headers = get_headers()
    response = requests.get(url, headers=headers)
    if response.status_code == 200:
        data = response.json()
        pretty_data = json.dumps(data, indent=4)
        print(pretty_data)
    else:
        print("Status Code:", response.status_code)
    return response.json()


def get_trigger_by_name(tn):
    trgs = get_triggers()
    for trg in trgs:
        if trg['triggerName'] == tn:
            return trg


def get_trigger_id_by_name(tn):
    trg = get_trigger_by_name(tn)
    if trg:
        return trg['triggerStrID']
    return '0'


def create_or_update_trigger(trigger_fn):
    url = api_v1_path + "/actions/ai"
    headers = get_headers()
    response = requests.post(url, json=trigger_fn, headers=headers)
    if response.status_code == 200:
        data = response.json()
        pretty_data = json.dumps(data, indent=4)
        print(pretty_data)
    else:
        print("Status Code:", response.status_code)
    return response.json()


trigger_function_template = {
    "triggerStrID": "0",
    "triggerID": 0,
    "triggerName": "demo-google-qp-search-regex",
    "triggerGroup": "demo",
    "triggerAction": "api-retrieval",
    "evalTriggerActions": [],
    "triggerRetrievals": []
}

trigger_function_eval = {
    "evalID": 0,
    "evalStrID": "0",
    "triggerID": 0,
    "triggerStrID": "0",
    "evalTriggerState": "filter",
    "evalResultsTriggerOn": "all-pass"
}

retrieval_template = {
    "retrievalStrID": "0",
    "retrievalID": 0,
    "retrievalName": "google-query-params",
    "retrievalGroup": "google-search",
    "retrievalItemInstruction": {
        "retrievalPlatform": "api",
        "retrievalPrompt": "",
        "retrievalPlatformGroups": "",
        "retrievalKeywords": "",
        "retrievalNegativeKeywords": "",
        "retrievalUsernames": "",
        "discordFilters": {
            "categoryTopic": "",
            "categoryName": "",
            "category": ""
        },
        "webFilters": {
            "routingGroup": "google-search",
            "lbStrategy": "",
            "maxRetries": 6,
            "backoffCoefficient": 2,
            "endpointRoutePath": "customsearch/v1?q={q}&cx=SEARCH_KEY&key=API_KEY",
            "endpointREST": "get",
            "payloadPreProcessing": "iterate",
            "regexPatterns": [
                "`https?://[^\\s<>\"]+|www\\.[^\\s<>\"]+`",
                "`og:([^\":]+)\":\\s*\"([^\"]+)`"
            ]
        }
    }
}

if __name__ == '__main__':
    with open('templates/trigger.json', 'r') as file:
        payload = json.load(file)

    # Add eval trigger
    trigger_function_eval['evalStrID'] = '1709068762300906000'
    trigger_function_template['evalTriggerActions'] = [trigger_function_eval]

    # Add retrieval for trigger
    retrieval_template['retrievalStrID'] = '1708579569890359000'
    trigger_function_template['triggerRetrievals'] = [retrieval_template]

    # Create or update trigger
    create_or_update_trigger(trigger_function_template)
