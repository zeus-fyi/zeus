import json

from examples.mockingbird.mockingbooks_py.workflows import create_wf


def create_google_regex_search_index_entities_wf(task_str_id, eval_str_id, agg_task_str_id, agg_eval_str_id):
    with open('mocks/workflow.json', 'r') as file:
        jdata = json.load(file)

    jdata['stepSize'] = 15
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
    # Add an eval to the workflow
    with open('../entities_triggers/mocks/eval_fn_create_entities.json', 'r') as file:
        agg_ef_data = json.load(file)

    jdata['evalsMap'] = {
        eval_str_id: ef_data,
        agg_eval_str_id: agg_ef_data
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
    # Add entities indexer trigger to agg output
    jdata['evalTasksMap'] = {
        task_str_id: {
            eval_str_id: True
        },
        agg_task_str_id: {
            agg_eval_str_id: True
        }
    }

    pretty_data = json.dumps(jdata, indent=4)
    print(pretty_data)
    create_wf(jdata)


if __name__ == '__main__':
    # create subcomponents first, if not already created

    task_str_id = '1708585225847375000'
    eval_str_id = '1708580672263367000'
    agg_task_str_id = '1708586569933699000'
    agg_eval_str_id = '1709417012941534000'
    create_google_regex_search_index_entities_wf(task_str_id, eval_str_id, agg_task_str_id, agg_eval_str_id)
