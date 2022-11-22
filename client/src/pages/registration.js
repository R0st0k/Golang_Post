import * as React from 'react';
import Head from "../components/head";
import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Stack from '@mui/material/Stack';
import TextField from '@mui/material/TextField';
import MenuItem from '@mui/material/MenuItem';
import {Typography} from "@mui/material";
import {Grid} from "@mui/material";
import CloseIcon from '@mui/icons-material/Close';
import Button from '@mui/material/Button';

import GetOrderIDDialog from "../components/getOrderIDDialog";

import "../styles/registration.css";

export default class Registration extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            type: "",
            sender: {
                name: "",
                surname: "",
                second_name: "",
                city: "",
                address: "",
                postcode: "",
                postcodes: []
            },
            receiver: {
                name: "",
                surname: "",
                second_name: "",
                city: "",
                address: "",
                postcode: "",
                postcodes: []
            },
            shape: {
                length: "",
                width: "",
                height: ""

            },
            weight: "",
            citysAndPostcodes: [
                {city: "Москва", postcodes: [345123, 125234, 677521]},
                {city: "Питер",  postcodes: [235232, 235619]}
            ],
            openOrderIDDialog: false
        }


        this.handleChange = this.handleChange.bind(this);
        this.handleChangeDefault= this.handleChangeDefault.bind(this);
        this.generateCitys = this.generateCitys.bind(this);
        this.handleSubmitButton = this.handleSubmitButton.bind(this);
        this.handleCloseDialog = this.handleCloseDialog.bind(this);
    }

    handleChange = typeOfInfo => event =>{
        const target = event.target;
        const value = target.value;
        const name = target.name;

        this.setState((prevState)=>({
            [typeOfInfo]: Object.assign({}, prevState[typeOfInfo], {
                [name]: value
            })
        }))
    }

    handleSubmitButton(event) {
        event.preventDefault()
        this.setState({
            openOrderIDDialog: true
        })
    }

    handleCloseDialog(){
        this.setState({
            openOrderIDDialog: false
        })
    }

    handleChangeDefault(event) {
        const target = event.target;
        const value = target.value;
        const name = target.name;

        this.setState({
                [name]: value
        })
    }

    handleChangeCity = typeOfInfo => event => {
        const target = event.target;
        const value = target.value;
        let object = this.state.citysAndPostcodes.find(element => element.city === value)
        this.setState((prevState)=>({
            [typeOfInfo]: Object.assign({}, prevState[typeOfInfo], {
                city: object.city,
                postcodes: object.postcodes
            })
        }))
    }

    generateCitys() {
        return this.state.citysAndPostcodes.map((data) => {
            return <MenuItem key={data.city} value={data.city}>{data.city}</MenuItem>
        });
    }




    render() {
        return (
            <>
                <Head/>
                <Box m={10}>
                <Card>
                    <CardContent>
                        <p>Регистрация нового отправления</p>
                        <form
                            onSubmit={this.handleSubmitButton}
                            autoComplete="off"
                        >
                            <Stack direction="row" spacing={1}>
                                <Typography>
                                    1. Выберите тип посылки
                                </Typography>
                                <Box sx={{
                                    width: 150
                                }}>
                                    <TextField
                                        required
                                        fullWidth
                                        select
                                        label="Тип"
                                        id="type"
                                        name="type"
                                        value={this.state.type}
                                        onChange={this.handleChangeDefault}
                                    >
                                        <MenuItem value={'Письмо'}>Письмо</MenuItem>
                                        <MenuItem value={'Посылка'}>Посылка</MenuItem>
                                        <MenuItem value={'Бандероль'}>Бандероль</MenuItem>
                                    </TextField>
                                </Box>
                            </Stack>
                            <Typography>
                                2. Отправитель
                            </Typography>
                            <Grid container spacing={1.5}>
                                <Grid item xs={3}>
                                    <TextField
                                        required
                                        name="surname"
                                        label="Фамилия"
                                        fullWidth
                                        value={this.state.sender.surname}
                                        onChange={this.handleChange('sender')}
                                    />
                                </Grid>
                                <Grid item xs={3}>
                                    <TextField
                                        required
                                        name="name"
                                        label="Имя"
                                        fullWidth
                                        value={this.state.sender.name}
                                        onChange={this.handleChange('sender')}
                                    />
                                </Grid>
                                <Grid item xs={3}>
                                    <TextField
                                        name="second_name"
                                        label="Отчество"
                                        fullWidth
                                        value={this.state.sender.second_name}
                                        onChange={this.handleChange('sender')}
                                    />
                                </Grid>
                                <Grid item xs={3}/>
                                <Grid item xs={3}>
                                    <TextField
                                        required
                                        name="city"
                                        label="Город"
                                        select
                                        fullWidth
                                        value={this.state.sender.city}
                                        onChange={this.handleChangeCity('sender')}
                                    >
                                        {this.generateCitys()}
                                    </TextField>
                                </Grid>
                                <Grid item xs={6}>
                                    <TextField
                                        label="Улица/дом/корпус/квартира"
                                        fullWidth
                                        name="address"
                                        value={this.state.sender.address}
                                        onChange={this.handleChange('sender')}
                                    />
                                </Grid>
                                <Grid item xs={3}>
                                    <TextField
                                        required
                                        name="postcode"
                                        select
                                        label="Индекс"
                                        fullWidth
                                        value={this.state.sender.postcode}
                                        onChange={this.handleChange('sender')}
                                    >
                                    {
                                        this.state.sender.postcodes.map((postcode) => (
                                            <MenuItem key={postcode} value={postcode}>{postcode}</MenuItem>
                                        ))
                                    }
                                    </TextField>
                                </Grid>
                            </Grid>
                            <Typography>
                                3. Получатель
                            </Typography>
                            <Grid container spacing={1.5}>
                                <Grid item xs={3}>
                                    <TextField
                                        required
                                        label="Фамилия"
                                        fullWidth
                                        name="surname"
                                        value={this.state.receiver.surname}
                                        onChange={this.handleChange('receiver')}
                                    />
                                </Grid>
                                <Grid item xs={3}>
                                    <TextField
                                        required
                                        label="Имя"
                                        name="name"
                                        fullWidth
                                        value={this.state.receiver.name}
                                        onChange={this.handleChange('receiver')}
                                    />
                                </Grid>
                                <Grid item xs={3}>
                                    <TextField
                                        label="Отчество"
                                        name="second_name"
                                        fullWidth
                                        value={this.state.receiver.second_name}
                                        onChange={this.handleChange('receiver')}
                                    />
                                </Grid>
                                <Grid item xs={3}/>
                                <Grid item xs={3}>
                                    <TextField
                                        required
                                        select
                                        label="Город"
                                        fullWidth
                                        value={this.state.receiver.city}
                                        onChange={this.handleChangeCity('receiver')}
                                    >
                                        {this.generateCitys()}
                                    </TextField>
                                </Grid>
                                <Grid item xs={6}>
                                    <TextField
                                        label="Улица/дом/корпус/квартира"
                                        name="address"
                                        fullWidth
                                        value={this.state.receiver.address}
                                        onChange={this.handleChange('receiver')}
                                    />
                                </Grid>
                                <Grid item xs={3}>
                                    <TextField
                                        required
                                        name="postcode"
                                        select
                                        label="Индекс"
                                        fullWidth
                                        value={this.state.receiver.postcode}
                                        onChange={this.handleChange('receiver')}
                                    >
                                    {
                                        this.state.receiver.postcodes.map((postcode) => (
                                            <MenuItem key={postcode} value={postcode}>{postcode}</MenuItem>
                                        ))
                                    }
                                    </TextField>
                                </Grid>
                            </Grid>
                        <Typography>
                            4. Характеристики
                        </Typography>
                        <Grid container spacing={1}>
                            <Grid item xs={3}>
                                <Stack direction="row">
                                    <TextField
                                        inputProps={{ inputMode: 'numeric', pattern: '^[0-9]+(\\.[0-9]+)?$' }}
                                        required
                                        name="length"
                                        label="Длина"
                                        fullWidth
                                        value={this.state.shape.length}
                                        onChange={this.handleChange('shape')}
                                    />
                                    <Typography mt={2} color={'#9F9B9B'}>
                                        мм
                                    </Typography>
                                    <Box mt={2}>
                                        <CloseIcon style={{color: '#9F9B9B'}}/>
                                    </Box>
                                </Stack>
                            </Grid>
                            <Grid item xs={3}>
                                <Stack direction="row">
                                    <TextField
                                        inputProps={{ inputMode: 'numeric', pattern: '^[0-9]+(\\.[0-9]+)?$' }}
                                        required
                                        name="width"
                                        label="Ширина"
                                        fullWidth
                                        value={this.state.shape.width}
                                        onChange={this.handleChange('shape')}
                                    />
                                    <Typography mt={2} color={'#9F9B9B'}>
                                        мм
                                    </Typography>
                                    <Box mt={2}>
                                        <CloseIcon style={{color: '#9F9B9B'}}/>
                                    </Box>
                                </Stack>
                            </Grid>
                            <Grid item xs={3}>
                                <Stack direction="row">
                                    <TextField
                                        inputProps={{ inputMode: 'numeric', pattern: '^[0-9]+(\\.[0-9]+)?$' }}
                                        required
                                        name="height"
                                        label="Высота"
                                        fullWidth
                                        value={this.state.shape.height}
                                        onChange={this.handleChange('shape')}
                                    />
                                    <Typography mt={2} color={'#9F9B9B'}>
                                        мм
                                    </Typography>
                                </Stack>
                            </Grid>
                            <Grid item xs={3}/>
                            <Grid item xs={3}>
                                <Stack direction="row">
                                    <TextField
                                        inputProps={{ inputMode: 'numeric', pattern: '^[0-9]+(\\.[0-9]+)?$' }}
                                        required
                                        name="weight"
                                        label="Вес"
                                        fullWidth
                                        value={this.state.weight}
                                        onChange={this.handleChangeDefault}
                                    />
                                    <Typography ml={1} mt={2} color={'#9F9B9B'}>
                                        г
                                    </Typography>
                                </Stack>
                            </Grid>
                        </Grid>
                            <Box className="submitButton">
                                <Stack spacing={1} direction="row">
                                    <Button variant="outlined" href={'/'}>
                                        Отмена
                                    </Button>
                                    <Button type="submit" variant="contained">
                                        Потвердить
                                    </Button>
                                </Stack>
                            </Box>
                    </form>
                    </CardContent>
                </Card>
                </Box>
                <GetOrderIDDialog open={this.state.openOrderIDDialog} onClose={this.handleCloseDialog}/>
            </>
        )
    }
}