import json

from examples.mockingbird.mockingbooks_py.workflows import create_wf

wf_model_task_template = {
    "taskStrID": "",
    "taskID": 0,
    "model": "",
    "taskType": "analysis",
    "temperature": 1.0,
    "marginBuffer": 0.5,
    "taskGroup": "group",
    "taskName": "name",
    "maxTokens": 0,
    "tokenOverflowStrategy": "deduce",
    "prompt": "",
    "cycleCount": 1,
    "evalFns": []
}

wf_model_agg_task_template = {
    "taskStrID": "",
    "taskID": 0,
    "model": "",
    "taskType": "aggregation",
    "temperature": 1.0,
    "marginBuffer": 0.5,
    "taskGroup": "group",
    "taskName": "name",
    "maxTokens": 0,
    "tokenOverflowStrategy": "deduce",
    "prompt": "",
    "cycleCount": 1,
    "evalFns": []
}
# Example Workflow Patterns


def create_analysis_only_wf(task_str_id):
    with open('templates/workflow.json', 'r') as file:
        jdata = json.load(file)

    jdata['workflowName'] = 'demo-analysis-only-workflow'
    jdata['workflowGroup'] = 'demo'

    jdata['stepSize'] = 5
    jdata['stepSizeUnit'] = 'minutes'

    # Add a task to the workflow
    wf_model_task_template['taskStrID'] = task_str_id
    jdata['models'][task_str_id] = wf_model_task_template

    pretty_data = json.dumps(jdata, indent=4)
    print(pretty_data)
    create_wf(jdata)


def create_analysis_ret_wf(task_str_id, retrieval_str_id):
    with open('templates/workflow.json', 'r') as file:
        jdata = json.load(file)

    jdata['workflowName'] = 'demo-analysis-ret-workflow'
    jdata['workflowGroup'] = 'demo'

    jdata['stepSize'] = 5
    jdata['stepSizeUnit'] = 'minutes'

    # Add a task to the workflow
    wf_model_task_template['taskStrID'] = task_str_id
    jdata['models'][task_str_id] = wf_model_task_template
    jdata['analysisRetrievalsMap'] = {
        task_str_id: {
            retrieval_str_id: True,
        }
    }
    pretty_data = json.dumps(jdata, indent=4)
    print(pretty_data)
    create_wf(jdata)


def create_analysis_eval_only_wf(task_str_id, eval_str_id):
    with open('templates/workflow.json', 'r') as file:
        jdata = json.load(file)

    jdata['workflowName'] = 'demo-analysis-eval-workflow'
    jdata['workflowGroup'] = 'demo'

    jdata['stepSize'] = 5
    jdata['stepSizeUnit'] = 'minutes'

    # Add a task to the workflow
    wf_model_task_template['taskStrID'] = task_str_id
    jdata['models'][task_str_id] = wf_model_task_template

    # Get eval fn template
    with open('templates/eval_fn.json', 'r') as file:
        ef_data = json.load(file)

    ef_data['evalStrID'] = eval_str_id
    # Add an eval to the workflow
    jdata['evalsMap'] = {
        eval_str_id: ef_data
    }
    jdata['evalTasksMap'] = {
        task_str_id: {
            eval_str_id: True
        }
    }
    pretty_data = json.dumps(jdata, indent=4)
    print(pretty_data)
    create_wf(jdata)


def create_agg_analysis_ret_wf(task_str_id, retrieval_id, agg_task_str_id):
    with open('templates/workflow.json', 'r') as file:
        jdata = json.load(file)

    jdata['workflowName'] = 'demo-agg-analysis-ret-workflow'
    jdata['workflowGroup'] = 'demo'

    jdata['stepSize'] = 5
    jdata['stepSizeUnit'] = 'minutes'

    # Add a task to the workflow
    wf_model_task_template['taskStrID'] = task_str_id
    jdata['models'][task_str_id] = wf_model_task_template

    # Add a retrieval to the workflow
    jdata['analysisRetrievalsMap'] = {
        task_str_id: {
            retrieval_id: True,
        }
    }

    # Add an aggregate task to the workflow
    wf_model_task_template['taskType'] = 'aggregate'
    wf_model_task_template['taskStrID'] = agg_task_str_id
    jdata['models'][agg_task_str_id] = wf_model_task_template

    jdata['aggregateSubTasksMap'] = {
        agg_task_str_id: {
            task_str_id: True
        }
    }

    pretty_data = json.dumps(jdata, indent=4)
    print(pretty_data)
    create_wf(jdata)


if __name__ == '__main__':
    # Set with your own values

    task_id = '1706393352827439000'
    ret_id = '1707094928902418000'
    agg_task_id = '1708910301810221000'
    eval_id = '1708049294114711000'

    create_analysis_only_wf(task_id)
    create_analysis_ret_wf(task_id, ret_id)
    create_agg_analysis_ret_wf(task_id, ret_id, agg_task_id)
    create_analysis_eval_only_wf(task_id, eval_id)
