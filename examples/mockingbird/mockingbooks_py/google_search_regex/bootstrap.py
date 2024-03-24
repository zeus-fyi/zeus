import json

from examples.mockingbird.mockingbooks_py.agg_tasks import create_agg_task
from examples.mockingbird.mockingbooks_py.analysis_tasks import create_analysis_task
from examples.mockingbird.mockingbooks_py.evals import create_or_update_eval
from examples.mockingbird.mockingbooks_py.retrievals import create_or_update_retrieval, get_retrieval_id_by_name
from examples.mockingbird.mockingbooks_py.schemas import create_or_update_schema, get_schema_by_name
from examples.mockingbird.mockingbooks_py.triggers import create_or_update_trigger, get_trigger_id_by_name


def schema_create():
    with open('mocks/schema.json', 'r') as file:
        data = json.load(file)
    create_or_update_schema(data)


def agg_schema_create():
    with open('mocks/agg_schema.json', 'r') as file:
        data = json.load(file)
    create_or_update_schema(data)


def retrieval_create():
    with open('mocks/retrieval.json', 'r') as file:
        data = json.load(file)
    create_or_update_retrieval(data)


def analysis_create():
    with open('mocks/analysis_task.json', 'r') as file:
        data = json.load(file)
    data['schemas'] = [get_schema_by_name('google-search-query-params')]
    create_analysis_task(data)


def agg_create():
    with open('mocks/agg_task.json', 'r') as file:
        data = json.load(file)
    data['schemas'] = [get_schema_by_name('google-search-query-params')]
    create_agg_task(data)


def agg_json_create():
    with open('mocks/agg_task_json_format.json', 'r') as file:
        data = json.load(file)
    data['schemas'] = [get_schema_by_name('google-search-results-agg')]
    create_agg_task(data)


def eval_fn_create():
    with open('mocks/eval_fn.json', 'r') as file:
        data = json.load(file)
    schema = get_schema_by_name('google-search-query-params')
    fields = schema['fields']
    for field in fields:
        if field['fieldName'] == 'q':
            field['evalMetrics'] = [
                {
                    "evalMetricID": 0,
                    "evalOperator": "length-less-than",
                    "evalState": "filter",
                    "evalExpectedResultState": "pass",
                    "evalMetricComparisonValues": {
                        "evalComparisonBoolean": False,
                        "evalComparisonNumber": 0,
                        "evalComparisonString": "100",
                        "evalComparisonInteger": 0
                    }
                },
            ]
    data['schemas'] = [schema]
    for tf in data['triggerFunctions']:
        tf['triggerStrID'] = get_trigger_id_by_name(tf['triggerName'])
    pretty_data = json.dumps(data, indent=4)
    print(pretty_data)
    create_or_update_eval(data)


def trigger_create():
    with open('mocks/trigger.json', 'r') as file:
        data = json.load(file)
    for tr in data['triggerRetrievals']:
        if tr['retrievalName'] == 'google-query-params':
            tr['retrievalStrID'] = get_retrieval_id_by_name('google-query-params')
    create_or_update_trigger(data)


if __name__ == '__main__':
    agg_json_create()
    schema_create()
    retrieval_create()
    trigger_create()
    eval_fn_create()
    analysis_create()
    agg_json_create()
