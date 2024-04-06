import json

from examples.mockingbird.mockingbooks_py.agg_tasks import create_agg_task
from examples.mockingbird.mockingbooks_py.analysis_tasks import create_analysis_task
from examples.mockingbird.mockingbooks_py.evals import create_or_update_eval, get_eval_id_by_name
from examples.mockingbird.mockingbooks_py.retrievals import create_or_update_retrieval, get_retrieval_by_name
from examples.mockingbird.mockingbooks_py.schemas import create_or_update_schema, get_schema_by_name
from examples.mockingbird.mockingbooks_py.triggers import create_or_update_trigger, get_trigger_id_by_name


def agg_create():
    with open('mocks/agg_task.json', 'r') as file:
        data = json.load(file)
    data['schemas'] = [get_schema_by_name('results-agg')]
    create_agg_task(data)


def retrieval_create():
    with open('mocks/retrieval.json', 'r') as file:
        data = json.load(file)
    create_or_update_retrieval(data)


def retrieval_biz_create():
    with open('mocks/retrieval_biz.json', 'r') as file:
        data = json.load(file)
    create_or_update_retrieval(data)


def agg_schema_create():
    with open('mocks/agg_schema.json', 'r') as file:
        data = json.load(file)
    create_or_update_schema(data)


def schema_create():
    with open('mocks/schema.json', 'r') as file:
        data = json.load(file)
    create_or_update_schema(data)


def analysis_create():
    with open('mocks/analysis_task.json', 'r') as file:
        data = json.load(file)
    data['schemas'] = [get_schema_by_name('linked-in-rapid-api-profile-qps')]
    create_analysis_task(data)


def trigger_create():
    with open('mocks/trigger.json', 'r') as file:
        data = json.load(file)
    data['triggerRetrievals'] = [get_retrieval_by_name('linkedin-profile')]
    data['evalTriggerAction'] = {
        "evalTriggerState": "filter",
        "evalResultsTriggerOn": "all-pass",
    }
    pretty_data = json.dumps(data, indent=4)
    print(pretty_data)
    create_or_update_trigger(data)


def eval_fn_create():
    with open('mocks/eval_fn.json', 'r') as file:
        data = json.load(file)
    schema = get_schema_by_name('linked-in-rapid-api-profile-qps')
    fields = schema['fields']
    for field in fields:
        if field['fieldName'] == 'linkedin_url':
            field['evalMetrics'] = [
                {
                    "evalMetricID": 0,
                    "evalOperator": "has-prefix",
                    "evalState": "filter",
                    "evalExpectedResultState": "pass",
                    "evalMetricComparisonValues": {
                        "evalComparisonString": "https://www.linkedin.com/in/",
                    }
                },
            ]
    data['schemas'] = [schema]
    data['triggerFunctions'] = [
        {
            "triggerStrID": get_trigger_id_by_name('linkedin-rapid-api-profile-search'),
            "triggerID": 0,
            "triggerName": "linkedin-rapid-api-profile-search",
            "triggerGroup": "linkedin",
            "triggerAction": "api-retrieval",
            "evalTriggerActions": [
                {
                    "evalID": 0,
                    "evalStrID": "0",
                    "triggerID": 0,
                    "triggerStrID": "0",
                    "evalTriggerState": "filter",
                    "evalResultsTriggerOn": "all-pass"
                }
            ]
        }
    ]
    data['evalStrID'] = get_eval_id_by_name('linkedin-rapid-api-profiles-qps')
    pretty_data = json.dumps(data, indent=4)
    print(pretty_data)
    create_or_update_eval(data)


if __name__ == '__main__':
    retrieval_biz_create()
    # schema_create()
    # analysis_create()
    # trigger_create()
    # eval_fn_create()
    # agg_schema_create()
    # agg_create()
    pass
