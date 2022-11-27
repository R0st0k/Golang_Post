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
import Autocomplete from '@mui/material/Autocomplete';

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
                middle_name: "",
                settlement: "",
                postcode: "",
                region: "",
                district: "",
                street: "",
                building: "",
                apartment: "",
                postcodes: []
            },
            receiver: {
                name: "",
                surname: "",
                middle_name: "",
                settlement: "",
                postcode: "",
                region: "",
                district: "",
                street: "",
                building: "",
                apartment: "",
                postcodes: []
            },
            size: {
                length: "",
                width: "",
                height: ""

            },
            weight: "",
            cityAndPostcodes: {
                "Санкт-Петербург": [
                    123001,
                    123002
                ],
                "Москва": [
                    124001
                ]
            },
            openOrderIDDialog: false
        }

        this.handleChange = this.handleChange.bind(this);
        this.handleChangeDefault= this.handleChangeDefault.bind(this);
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

    handleChangeCity = typeOfInfo => (event, value) => {
        const postcodes = this.state.cityAndPostcodes[value];
        if (typeof postcodes === 'undefined') return;
        this.setState((prevState)=>({
            [typeOfInfo]: Object.assign({}, prevState[typeOfInfo], {
                settlement: value,
                postcodes: postcodes
            })
        }))
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
                                        name="middle_name"
                                        label="Отчество"
                                        fullWidth
                                        value={this.state.sender.middle_name}
                                        onChange={this.handleChange('sender')}
                                    />
                                </Grid>
                                <Grid item xs={3}/>
                                <Grid item xs={3}>
                                    <Autocomplete
                                        name="city"
                                        renderInput={(params => <TextField
                                            {...params}
                                            label={"Поселение"}
                                        />)}
                                        onChange={this.handleChangeCity('sender')}
                                        options={Object.keys(this.state.cityAndPostcodes)}
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
                                <Grid item xs={3}>
                                    <TextField
                                        label="Регион"
                                        fullWidth
                                        name="region"
                                        value={this.state.sender.region}
                                        onChange={this.handleChange('sender')}
                                    />
                                </Grid>
                                <Grid item xs={3}>
                                    <TextField
                                        label="Округ"
                                        fullWidth
                                        name="district"
                                        value={this.state.sender.district}
                                        onChange={this.handleChange('sender')}
                                    />
                                </Grid>
                                <Grid item xs={3}>
                                    <TextField
                                        label="Улица"
                                        fullWidth
                                        name="street"
                                        value={this.state.sender.street}
                                        onChange={this.handleChange('sender')}
                                    />
                                </Grid>
                                <Grid item xs={3}>
                                    <TextField
                                        label="Корпус"
                                        fullWidth
                                        name="building"
                                        value={this.state.sender.building}
                                        onChange={this.handleChange('sender')}
                                        inputProps={{ inputMode: 'numeric', pattern: '^[0-9]+$' }}
                                    />
                                </Grid>
                                <Grid item xs={3}>
                                    <TextField
                                        label="Квартира"
                                        fullWidth
                                        name="apartment"
                                        value={this.state.sender.apartment}
                                        onChange={this.handleChange('sender')}
                                        inputProps={{ inputMode: 'numeric', pattern: '^[0-9]+$' }}
                                    />
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
                                        name="middle_name"
                                        fullWidth
                                        value={this.state.receiver.middle_name}
                                        onChange={this.handleChange('receiver')}
                                    />
                                </Grid>
                                <Grid item xs={3}/>
                                <Grid item xs={3}>
                                    <Autocomplete
                                        name="city"
                                        renderInput={(params => <TextField
                                            {...params}
                                            label={"Поселение"}
                                        />)}
                                        onChange={this.handleChangeCity('receiver')}
                                        options={Object.keys(this.state.cityAndPostcodes)}
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
                                <Grid item xs={3}>
                                    <TextField
                                        label="Регион"
                                        fullWidth
                                        name="region"
                                        value={this.state.receiver.region}
                                        onChange={this.handleChange('receiver')}
                                    />
                                </Grid>
                                <Grid item xs={3}>
                                    <TextField
                                        label="Округ"
                                        fullWidth
                                        name="district"
                                        value={this.state.receiver.district}
                                        onChange={this.handleChange('receiver')}
                                    />
                                </Grid>
                                <Grid item xs={3}>
                                    <TextField
                                        label="Улица"
                                        fullWidth
                                        name="street"
                                        value={this.state.receiver.street}
                                        onChange={this.handleChange('receiver')}
                                    />
                                </Grid>
                                <Grid item xs={3}>
                                    <TextField
                                        label="Корпус"
                                        fullWidth
                                        name="building"
                                        value={this.state.receiver.building}
                                        onChange={this.handleChange('receiver')}
                                        inputProps={{ inputMode: 'numeric', pattern: '^[0-9]+$' }}
                                    />
                                </Grid>
                                <Grid item xs={3}>
                                    <TextField
                                        label="Квартира"
                                        fullWidth
                                        name="apartment"
                                        value={this.state.receiver.apartment}
                                        onChange={this.handleChange('receiver')}
                                        inputProps={{ inputMode: 'numeric', pattern: '^[0-9]+$' }}
                                    />
                                </Grid>
                            </Grid>
                        <Typography>
                            4. Характеристики
                        </Typography>
                        <Grid container spacing={1}>
                            <Grid item xs={3}>
                                <Stack direction="row">
                                    <TextField
                                        inputProps={{ inputMode: 'numeric', pattern: '^[0-9]{2,}$' }}
                                        required
                                        name="length"
                                        label="Длина"
                                        fullWidth
                                        value={this.state.size.length}
                                        onChange={this.handleChange('size')}
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
                                        inputProps={{ inputMode: 'numeric', pattern: '^[0-9]{2,}$' }}
                                        required
                                        name="width"
                                        label="Ширина"
                                        fullWidth
                                        value={this.state.size.width}
                                        onChange={this.handleChange('size')}
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
                                        inputProps={{ inputMode: 'numeric', pattern: '^[0-9]{2,}$' }}
                                        required
                                        name="height"
                                        label="Высота"
                                        fullWidth
                                        value={this.state.size.height}
                                        onChange={this.handleChange('size')}
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
                                        inputProps={{ inputMode: 'numeric', pattern: '^[0-9]{2,}$' }}
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