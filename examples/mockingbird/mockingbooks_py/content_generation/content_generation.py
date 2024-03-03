import json

from examples.mockingbird.mockingbooks_py.analysis_tasks import create_analysis_task
from examples.mockingbird.mockingbooks_py.evals import create_or_update_eval
from examples.mockingbird.mockingbooks_py.retrievals import create_or_update_retrieval
from examples.mockingbird.mockingbooks_py.schemas import create_or_update_schema
from examples.mockingbird.mockingbooks_py.triggers import create_or_update_trigger
from examples.mockingbird.mockingbooks_py.workflows import start_or_schedule_wf, create_wf

wf_model_task_template = {
    "taskStrID": "",
    "taskID": 0,
    "model": "gpt-3.5-turbo-0125",
    "taskType": "analysis",
    "temperature": 1.0,
    "marginBuffer": 0.5,
    "taskGroup": "group",
    "taskName": "name",
    "responseFormat": "json",
    "maxTokens": 0,
    "tokenOverflowStrategy": "deduce",
    "prompt": "",
    "cycleCount": 1,
    "evalFns": []
}


def create_schema():
    with open('mocks/schema.json', 'r') as file:
        schema_data = json.load(file)
    create_or_update_schema(schema_data)


def create_api_call():
    with open('mocks/api_call.json', 'r') as file:
        api_call_json = json.load(file)
    create_or_update_retrieval(api_call_json)


def create_llm_tweet_retrieval():
    with open('mocks/retrieval_llm_tweets.json', 'r') as file:
        ret_twitter_json = json.load(file)
    create_or_update_retrieval(ret_twitter_json)


def create_ef_tweet():
    with open('mocks/eval_fn_trigger.json', 'r') as file:
        eval_fn_pl = json.load(file)
    create_or_update_eval(eval_fn_pl)


def create_trigger_tweet():
    with open('mocks/trigger_tweet.json', 'r') as file:
        tlt = json.load(file)
    create_or_update_trigger(tlt)


def create_like_tweets_wf():
    with open('mocks/workflow.json', 'r') as file:
        wf_json = json.load(file)
    create_wf(wf_json)


def create_tweet_analysis_task():
    with open('mocks/tweet_analysis.json', 'r') as file:
        at = json.load(file)
    create_analysis_task(at)


def create_tweet_cg_wf():
    with open('mocks/workflow.json', 'r') as file:
        jdata = json.load(file)

    # Add a task to the workflow
    task_str_id = '1709494294817279000'
    wf_model_task_template['taskStrID'] = task_str_id
    jdata['models'][task_str_id] = wf_model_task_template

    # Get eval fn template
    with open('mocks/eval_fn_trigger.json', 'r') as file:
        ef_data = json.load(file)

    # Add an eval to the workflow
    jdata['evalsMap'] = {
        ef_data['evalStrID']: ef_data
    }
    jdata['evalTasksMap'] = {
        task_str_id: {
            ef_data['evalStrID']: True
        }
    }
    pretty_data = json.dumps(jdata, indent=4)
    print(pretty_data)
    create_wf(jdata)


def create_tweet_like_wf():
    # step 1. create schema def
    create_schema()

    # step 2. create api call definition
    create_api_call()

    # step 3. create a trigger
    create_trigger_tweet()

    # step 4. create eval fn that triggers a tweet like POST request
    create_ef_tweet()

    # step 5. create extract tweets task
    create_tweet_analysis_task()

    # step 6. create a workflow that uses the eval fn
    create_tweet_cg_wf()

    # optionally. replace the json id str values with the real ones to update the mocks via the same functions


def start_wf():
    # content source
    with open('content/example.txt', 'r') as file:
        content_input = file.read()

    # then exec workflow
    with open('mocks/wf_exec.json', 'r') as file:
        wf_exec = json.load(file)

    wf_item = {
        "workflowName": "tweet_content_writer_wf",
    }
    wf_exec['workflows'] = [wf_item]

    tmp = {'tweet_create_content': {'replacePrompt': content_input}}
    wf_exec['taskOverrides'] = tmp

    pretty_data = json.dumps(wf_exec, indent=4)
    print(pretty_data)

    start_or_schedule_wf(wf_exec)


if __name__ == '__main__':
    # create_schema()

    start_wf()
    # lookup wf run id you get from the response
    # you can poll this for the status of the workflow run

    # you can override the mocks or just use the json file definition as is.
    # This unlocks easy GitOps for your ML production CI flow using the .json files as the source of truth.

    # e.g. override the workflow json with custom values

    # unix_start_time_sec = 1707883380
    # run_time_duration = 1
    # run_time_unit = "hours"
