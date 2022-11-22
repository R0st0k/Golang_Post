import * as React from 'react';
import Head from "../components/head";
import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';

import CustomAccordion from "../components/customAccordion";
import AdvancedSearchSendings from "../components/advancedSearchSendings";
import SendingsTable from "../components/sendingsTable";


export default class SendingsInfo extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            orderID: "",
            advanced_search: {
                type: "",
                date: "",
                status: "",
                departure_city: "",
                arrival_city: "",
                weight: "",
                length: "",
                width: "",
                height: ""
            }
        }
        this.handleChange = this.handleChange.bind(this);
        this.handleChangeAdvanceSearch = this.handleChangeAdvanceSearch.bind(this);
        this.handleSubmitButton = this.handleSubmitButton.bind(this);
    }

    handleChange(event){
        const target = event.target;
        const value = target.value;
        const name = target.name;
        this.setState({
            [name]: value
        })
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
                                        name="orderID"
                                        label="order-id"
                                        value={this.state.orderID}
                                        onChange={this.handleChange}
                                    />
                                    <Button style={{marginLeft: "3%"}} type={"submit"} variant={"contained"}>
                                        Показать
                                    </Button>
                                </Box>
                                <CustomAccordion
                                    label={"Расширенный поиск"}
                                    element={<AdvancedSearchSendings values={this.state.advanced_search} onChange={this.handleChangeAdvanceSearch}/>}
                                />
                            </form>
                        </CardContent>
                    </Card>
                </Box>
                <SendingsTable/>
            </>
        )
    }
}