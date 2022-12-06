import * as React from 'react';
import Head from "../components/head";
import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import axios from 'axios';

import CustomAccordion from "../components/customAccordion";
import AdvancedSearchSendings from "../components/advancedSearchSendings";
import SendingsTable from "../components/sendingsTable";

export default class SendingsInfo extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            order_id: "",
            filter:  {},
            advanced_search: {
                type: [],
                date_start: "",
                date_finish: "",
                status: [],
                sender_settlement: "",
                receiver_settlement: "",
                weight: "",
                length: "",
                width: "",
                height: ""
            },
            data: [],
            total: 0,
            page: 0,
            rowsPerPage: 5,
            order: "asc",
            orderBy: ""
        }
        this.handleChange = this.handleChange.bind(this);
        this.handleChangeAdvanceSearch = this.handleChangeAdvanceSearch.bind(this);
        this.handleSubmitButton = this.handleSubmitButton.bind(this);
        this.handleChangeTable = this.handleChangeTable.bind(this);
        this.refreshTable = this.refreshTable.bind(this);
        this.handleChangeRowsPerPage = this.handleChangeRowsPerPage.bind(this);
    }

    componentDidMount() {
        axios.get('http://localhost:8080/api/v1/sending_filter', {
            params:{
                page: this.state.page + 1,
                elems_on_page: this.state.rowsPerPage
            }
        })
            .then(
                (response) => {
                    this.setState({
                        data: response.data.result,
                        total: response.data.total
                    })
                }
            )
    }

    handleChange(event){
        const target = event.target;
        const value = target.value;
        const name = target.name;
        this.setState({
            [name]: value
        });
    }

    handleChangeTable(target){
        const name = target.name;
        const value = target.value;
        this.setState({
            [name]: value
        }, () =>{
            this.refreshTable();
        })
    }

    handleChangeRowsPerPage(value){
        let page = this.state.page;
        const total = this.state.total;
        if(value * page >= total){
            page = Math.ceil(total / value) - 1;
        }
        this.setState({
            page: page,
            rowsPerPage: value
        }, () => {
            this.refreshTable();
        })
    }

    refreshTable(){
        let params = {};
        const filter = this.state.filter;
        for(let key in filter){
            if(filter[key] !== ""){
                params[key] = filter[key];
            }
        }
        params["page"] = this.state.page + 1;
        params["elems_on_page"] = this.state.rowsPerPage;
        if(this.state.order_id !== ""){
            params['order_id'] = this.state.order_id;
        }
        if(this.state.orderBy !== ""){
            params['sort_field'] = this.state.orderBy;
            if(this.state.order !== ""){
                params['sort_type'] = this.state.order;
            }
        }
        axios.get('http://localhost:8080/api/v1/sending_filter', {
            params
        })
            .then(
                (response) => {
                    this.setState({
                        data: response.data.result,
                        total: response.data.total,
                    })
                }
            )
    }

    handleChangeAdvanceSearch(event){
        const target = event.target;
        const value = target.value;
        const name = target.name;
        this.setState((prevState)=>({
            advanced_search: Object.assign({}, prevState.advanced_search, {
                [name]: value
            })
        }))
    }

    handleSubmitButton(event){
        event.preventDefault();
        this.setState({
            filter: this.state.advanced_search,
            page: 0
        },  () => {
            this.refreshTable();
        })
    }

    render() {
        return (
            <>
                <Head />
                <Box ml={"15%"} mr={"15%"} mt={10}>
                    <Card>
                        <CardContent>
                            <form
                                onSubmit={this.handleSubmitButton}
                            >
                                <Box textAlign={"center"} mt={1}>
                                    Введите данные отправления
                                </Box>
                                <Box textAlign={"center"}  mt={1}>
                                    <TextField
                                        name="order_id"
                                        label="order-id"
                                        value={this.state.order_id}
                                        onChange={this.handleChange}
                                    />
                                    <Button style={{marginLeft: "3%"}} type={"submit"} variant={"contained"}>
                                        Показать
                                    </Button>
                                </Box>
                                <CustomAccordion
                                    label={"Расширенный поиск"}
                                    element={
                                        <AdvancedSearchSendings
                                            values={this.state.advanced_search}
                                            onChange={this.handleChangeAdvanceSearch}
                                    />}
                                />
                            </form>
                        </CardContent>
                    </Card>
                </Box>
                <SendingsTable
                    total={this.state.total}
                    data={this.state.data}
                    page={this.state.page}
                    rowsPerPage={this.state.rowsPerPage}
                    order={this.state.order}
                    orderBy={this.state.orderBy}
                    handleChangeTable={this.handleChangeTable}
                    onChangeRowsPerPage={this.handleChangeRowsPerPage}

                />
            </>
        )
    }
}