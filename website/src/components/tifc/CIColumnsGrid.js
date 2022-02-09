import { CIErrorTable } from "./CIErrorTable";

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
        flex: 10,
        align: 'left',
        editable: false,
        filterable: false,
        pinnable: true,
    },
    {
        field: 'case_type',
        headerName: 'Failure Type',
        headerAlign: 'left',
        flex: 10,
        align: 'left',
        editable: false,
    },
    {
        field: 'case_status',
        headerName: 'Failure Status',
        headerAlign: 'left',
        flex: 10,
        align: 'left',
        editable: false,
    },
    {
        field: 'test_case_info',
        headerName: 'Test Case Info',
        headerAlign: 'left',
        flex: 30,
        valueGetter: function(params){
            return 'TestSuiteName: ' + params.row.test_suite_name + '<br/>'
                + 'TestCaseName: ' + params.row.test_case_name + '<br/>'
                + 'TestClassName: ' + params.row.test_class_name;
        },
        sortable: false,
        align: 'left',
        editable: false,
    },
    {
        field: 'first_introducer',
        headerName: 'First Introducer',
        headerAlign: 'left',
        flex: 10,
        align: 'left',
        editable: false,
    },
    {   
        field: 'resource_cost',
        headerName: 'Resource Cost',
        headerAlign: 'left',
        flex: 10,
        align: 'left',
        editable: false,
        filterable: false,
    },
    {
        field: 'first_seen',
        headerName: 'First Seen',
        headerAlign: 'left',
        flex: 30,
        valueGetter: function(params){
            return 'PrID: ' + params.row.first_seen.pull_request + '<br/>'
                + 'CommitID: ' + params.row.first_seen.commit_id + '<br/>'
                + 'Author: ' + params.row.first_seen.author;
        },
        sortable: false,
        align: 'left',
        editable: false,
        filterable: false,
    },
    {
        field: 'last_seen',
        headerName: 'Last Seen',
        headerAlign: 'left',
        flex: 30,
        valueGetter: function(params){
            return 'PrID: ' + params.row.last_seen.pull_request + '<br/>'
                + 'CommitID: ' + params.row.last_seen.commit_id + '<br/>'
                + 'Author: ' + params.row.last_seen.author;
        },
        sortable: false,
        align: 'left',
        editable: false,
        filterable: false,
    },
    {
        field: 'recent_runs',
        headerName: 'Recent Runs',
        headerAlign: 'left',
        flex: 30,
        renderCell: function(params){
            return <CIErrorTable data={params.row.recent_runs} />
        },
        sortable: false,
        align: 'left',
        editable: false,
        filterable: false,
    },
];

export default CIColumnsGrid;