import { CIRecentRuns } from "./CIRecentRuns";

const CIColumnsGrid = [
    {
        field: 'id',
        headerName: 'ID',
        headerAlign: 'left',
        hide: true,
        align: 'left',
        editable: false,
        filterable: false,
    },
    { 
        field: 'failed_count',
        headerName: 'Failed Count',
        headerAlign: 'left',
        type: 'number',
        align: 'left',
        editable: false,
        filterable: false,
        pinnable: true,
        minWidth: 120,
        resizable: true,
    },
    {
        field: 'case_type',
        headerName: 'Failure Type',
        headerAlign: 'left',
        align: 'left',
        editable: false,
        minWidth: 120,
    },
    {
        field: 'case_status',
        headerName: 'Failure Status',
        headerAlign: 'left',
        align: 'left',
        editable: false,
        minWidth: 120,
    },
    {
        field: 'test_case_info',
        headerName: 'Test Case Info',
        headerAlign: 'left',
        valueGetter: function(params) {
            // for download
            return `TestSuiteName: ${params.row.test_suite_name || ''} | TestCaseName: ${params.row.test_case_name || ''} | TestClassName: ${params.row.test_class_name || ''}`;
        },
        renderCell: function(params) {
            return <p>
                TestSuiteName: {params.row.test_suite_name} <br/>
                TestCaseName: {params.row.test_case_name} <br/>
                TestClassName: {params.row.test_class_name} 
            </p>;
        },
        sortable: false,
        align: 'left',
        editable: false,
        filterable: true,
        minWidth: 400,
    },
    {
        field: 'first_introducer',
        headerName: 'First Introducer',
        headerAlign: 'left',
        align: 'left',
        editable: false,
        minWidth: 130,
    },
    {
        field: 'first_seen',
        headerName: 'First Seen',
        headerAlign: 'left',
        valueGetter: function(params) {
            // for download
            return JSON.stringify(params.row.first_seen, null, 4);
        },
        renderCell: function(params) {
            return <p>
                PRID: {params.row.first_seen.pull_request} <br/>
                CommitID: {params.row.first_seen.commit_id} <br/>
                Author: {params.row.first_seen.author} 
            </p>;
        },
        sortable: false,
        align: 'left',
        editable: false,
        filterable: false,
        minWidth: 400,
    },
    {
        field: 'last_seen',
        headerName: 'Last Seen',
        headerAlign: 'left',
        valueGetter: function(params) {
            // for download
            return JSON.stringify(params.row.last_seen, null, 4);
        },
        renderCell: function(params) {
            return <p>
                PRID: {params.row.last_seen.pull_request} <br/>
                CommitID: {params.row.last_seen.commit_id} <br/>
                Author: {params.row.last_seen.author} 
            </p>;
        },
        sortable: false,
        align: 'left',
        editable: false,
        filterable: false,
        minWidth: 400,
    },
    {
        field: 'may_introduced_by',
        headerName: 'May Introduced By',
        headerAlign: 'left',
        valueGetter: function(params) {
            // for download
            return JSON.stringify(params.row.may_introduced_by, null, 4);
        },
        renderCell: function(params) {
            // for page show
            return <p>
                PRID: {params.row.may_introduced_by.pull_request} <br/>
                CommitID: {params.row.may_introduced_by.commit_id} <br/>
                Author: {params.row.may_introduced_by.author} 
            </p>;
        },
        sortable: false,
        align: 'left',
        editable: false,
        filterable: false,
        minWidth: 400,
    },
    {   
        field: 'resource_cost',
        headerName: 'Resource Cost',
        headerAlign: 'left',
        align: 'left',
        editable: false,
        filterable: false,
        minWidth: 130,
    },
    {   
        field: 'failure_age',
        headerName: 'Failure Age',
        headerAlign: 'left',
        align: 'left',
        editable: false,
        filterable: false,
        minWidth: 130,
    },
    {
        field: 'recent_runs',
        headerName: 'Trace History',
        headerAlign: 'left',
        valueGetter: function(params) {
            // for download
            return JSON.stringify(params.row.recent_runs, null, 4);
        },
        renderCell: function(params){
            return <CIRecentRuns data={params.row.recent_runs} />
        },
        sortable: false,
        align: 'left',
        editable: false,
        filterable: false,
        minWidth: 120,
    },
];

export default CIColumnsGrid;
