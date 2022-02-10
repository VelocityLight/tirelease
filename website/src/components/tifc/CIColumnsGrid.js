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
        valueGetter: function(params){
            return 'TestSuiteName: ' + params.row.test_suite_name + '<br/>'
                + 'TestCaseName: ' + params.row.test_case_name + '<br/>'
                + 'TestClassName: ' + params.row.test_class_name;
        },
        sortable: false,
        align: 'left',
        editable: false,
        filterable: true,
        minWidth: 300,
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
        field: 'resource_cost',
        headerName: 'Resource Cost',
        headerAlign: 'left',
        align: 'left',
        editable: false,
        filterable: false,
        minWidth: 130,
    },
    {
        field: 'first_seen',
        headerName: 'First Seen',
        headerAlign: 'left',
        valueGetter: function(params){
            return 'PRID: ' + params.row.first_seen.pull_request + '<br/>'
                + 'CommitID: ' + params.row.first_seen.commit_id + '<br/>'
                + 'Author: ' + params.row.first_seen.author;
        },
        sortable: false,
        align: 'left',
        editable: false,
        filterable: false,
        minWidth: 200,
    },
    {
        field: 'last_seen',
        headerName: 'Last Seen',
        headerAlign: 'left',
        valueGetter: function(params){
            return 'PRID: ' + params.row.last_seen.pull_request + '<br/>'
                + 'CommitID: ' + params.row.last_seen.commit_id + '<br/>'
                + 'Author: ' + params.row.last_seen.author;
        },
        sortable: false,
        align: 'left',
        editable: false,
        filterable: false,
        minWidth: 200,
    },
    {
        field: 'recent_runs',
        headerName: 'Trace Logs',
        headerAlign: 'left',
        renderCell: function(params){
            return <CIRecentRuns data={params.row.recent_runs} />
        },
        sortable: false,
        align: 'left',
        editable: false,
        filterable: false,
    },
];

export default CIColumnsGrid;
