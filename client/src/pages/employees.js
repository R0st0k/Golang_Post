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
                birthday_start: "",
                birthday_end: "",
                sex: "",
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
            //this.refreshTable();
        })
    }

    handleChangeTable(target){
        const name = target.name;
        const value = target.value;
        this.setState({
            [name]: value
        }, () =>{
            //this.refreshTable();
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
            //this.refreshTable();
        })
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