import json

from examples.mockingbird.mockingbooks_py.analysis_tasks import get_task_id_by_name
from examples.mockingbird.mockingbooks_py.evals import get_eval_id_by_name
from examples.mockingbird.mockingbooks_py.workflows import create_wf


def create_linkedin_rapid_api_wf(task_str_id, eval_str_id, agg_task_str_id, agg_eval_str_id=None):
    with open('mocks/workflow.json', 'r') as file:
        jdata = json.load(file)

    jdata['stepSize'] = 30
    jdata['stepSizeUnit'] = 'minutes'

    # Add a task to the workflow
    with open('mocks/analysis_task.json', 'r') as file:
        wf_analysis_task = json.load(file)

    wf_analysis_task['taskStrID'] = task_str_id
    wf_analysis_task['taskType'] = 'analysis'
    wf_analysis_task['modelType'] = 'gpt-4-0125-preview'
    jdata['models'][task_str_id] = wf_analysis_task

    # Get eval fn template
    with open('mocks/eval_fn.json', 'r') as file:
        ef_data = json.load(file)

    ef_data['evalStrID'] = eval_str_id

    jdata['evalsMap'] = {
        eval_str_id: ef_data,
    }

    # Add an aggregate task to the workflow
    with open('mocks/agg_task.json', 'r') as file:
        wf_agg_task = json.load(file)

    wf_agg_task['taskStrID'] = agg_task_str_id
    wf_agg_task['taskType'] = 'aggregation'
    jdata['models'][agg_task_str_id] = wf_agg_task
    jdata['aggregateSubTasksMap'] = {
        agg_task_str_id: {
            task_str_id: True
        }
    }
    jdata['evalTasksMap'] = {
        task_str_id: {
            eval_str_id: True
        },
    }

    pretty_data = json.dumps(jdata, indent=4)
    print(pretty_data)
    create_wf(jdata)


if __name__ == '__main__':
    task_id = get_task_id_by_name('linkedin-profiles-rapid-api-qps')
    eval_id = get_eval_id_by_name('linkedin-rapid-api-profiles-qps')
    agg_task_id = get_task_id_by_name('linkedin-search-summary')
    create_linkedin_rapid_api_wf(task_id, eval_id, agg_task_id)
