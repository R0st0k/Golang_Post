import * as React from 'react';
import Head from "../components/head";
import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import {Grid} from "@mui/material";
import axios from 'axios';

import CustomAccordion from "../components/customAccordion";
import AdvancedSearchEmployees from "../components/advancedSearchEmployees";
import EmployeesTable from "../components/employeesTable";
import qs from "qs";


export default class Employees extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            surname: "",
            name: "",
            middle_name: "",
            advanced_search: {
                settlement: "",
                postcode: "",
                position: "",
                birth_date_start: "",
                birth_date_end: "",
                gender: "",
                phone_number: ""
            },
            data: [],
            filter: {},
            total: 1,
            page: 0,
            rowsPerPage: 5,
            order: "asc",
            orderBy: ""
        }
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmitButton = this.handleSubmitButton.bind(this);
        this.handleChangeAdvanceSearch = this.handleChangeAdvanceSearch.bind(this);
        this.handleChangeTable = this.handleChangeTable.bind(this);
        this.handleChangeRowsPerPage = this.handleChangeRowsPerPage.bind(this);
    }

    componentDidMount() {
        axios.get('http://localhost:8080/api/v1/employee_filter', {
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
        if(this.state.surname !== ""){
            params['surname'] = this.state.surname;
        }
        if(this.state.name !== ""){
            params['name'] = this.state.name;
        }
        if(this.state.middle_name !== ""){
            params['middle_name'] = this.state.middle_name;
        }
        if(this.state.orderBy !== ""){
            params['sort_field'] = this.state.orderBy;
            if(this.state.order !== ""){
                params['sort_type'] = this.state.order;
            }
        }
        const api = axios.create({
            paramsSerializer: {
                serialize: (params) => qs.stringify(params, {arrayFormat: 'repeat'})
            }
        });
        api.get('http://localhost:8080/api/v1/employee_filter',{
            params: params
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

    render(){
        return(
            <>
                <Head/>
                <Box ml={"15%"} mr={"15%"} mt={10}>
                    <Card>
                        <CardContent>
                            <form
                                onSubmit={this.handleSubmitButton}
                            >
                                <Box textAlign={"center"} mt={1}>
                                    Введите данные сотрудника
                                </Box>
                                <Grid container spacing={2}>
                                    <Grid item xs={3}>
                                        <TextField
                                            fullWidth
                                            name="surname"
                                            label="Фамилия"
                                            value={this.state.surname}
                                            onChange={this.handleChange}
                                        />
                                    </Grid>
                                    <Grid item xs={3}>
                                        <TextField
                                            fullWidth
                                            name="name"
                                            label="Имя"
                                            value={this.state.name}
                                            onChange={this.handleChange}
                                        />
                                    </Grid>
                                    <Grid item xs={3}>
                                        <TextField
                                            fullWidth
                                            name="middle_name"
                                            label="Фамилия"
                                            value={this.state.middle_name}
                                            onChange={this.handleChange}
                                        />
                                    </Grid>
                                    <Grid item xs={3}>
                                        <Button style={{marginLeft: "3%"}} type={"submit"} variant={"contained"}>
                                            Показать
                                        </Button>
                                    </Grid>
                                </Grid>
                                <CustomAccordion
                                    label={"Расширенный поиск"}
                                    element={
                                        <AdvancedSearchEmployees
                                            values={this.state.advanced_search}
                                            onChange={this.handleChangeAdvanceSearch}
                                        />}
                                />
                            </form>
                        </CardContent>
                    </Card>
                </Box>
                <EmployeesTable
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