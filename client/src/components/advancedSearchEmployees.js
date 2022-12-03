import {Grid, Typography} from "@mui/material";
import TextField from "@mui/material/TextField";
import MenuItem from "@mui/material/MenuItem";
import Stack from "@mui/material/Stack";
import Box from "@mui/material/Box";
import CloseIcon from "@mui/icons-material/Close";
import { AdapterMoment } from '@mui/x-date-pickers/AdapterMoment';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DesktopDatePicker } from '@mui/x-date-pickers/DesktopDatePicker';
import * as React from "react";
import {useEffect, useState} from "react";
import axios from "axios";
import Autocomplete from "@mui/material/Autocomplete";
import MuiPhoneNumber from 'material-ui-phone-number-2'

export default function AdvancedSearchEmployees(props){
    const [cityAndPostcodes, setCityAndPostcodes] = useState([]);
    const [postcodes, setPostcodes] = useState([]);

    useEffect(() => {
        axios.get('http://localhost:8080/api/v1/postcodes_by_settlement')
            .then(
                (response) => {
                    setCityAndPostcodes(response.data);
                }
            )
    }, []);

    /*const handleChangeCity = name => (event, value) => {
        let my_event = {};
        const target = {name: name, value: value};
        my_event.target = target;
        props.onChange(my_event);
    }*/

    const handleChangeCity = (event, value) => {
        let my_event = {};
        const target = {name: "settlement", value: value};
        my_event.target = target;
        props.onChange(my_event);
        const postcodes = cityAndPostcodes[value];
        if (typeof postcodes === 'undefined') return;
        setPostcodes(postcodes);
    }

    const handleDateStartChange = value =>{
        let event = {};
        const target = {name: "birthday_start", value: value.format('YYYY-M-DD')};
        event.target = target;
        props.onChange(event);
    }

    const handleDateEndChange = value =>{
        let event = {};
        const target = {name: "birthday_end", value: value.format('YYYY-M-DD')};
        event.target = target;
        props.onChange(event);
    }

    const handleNumberPhoneChange = value =>{
        let event = {};
        const target = {name: "phone_number", value: value};
        event.target = target;
        props.onChange(event);
    }


    return (
        <div>
            <Grid container spacing={2}>
                <Grid item xs={4.5}>
                    <Autocomplete
                        fullWidth
                        name="settlement"
                        renderInput={(params => <TextField
                            {...params}
                            label={"Населенный пункт"}
                        />)}
                        onChange={handleChangeCity}
                        options={Object.keys(cityAndPostcodes)}
                    />
                </Grid>
                <Grid item xs={3}>
                    <TextField
                        name="postcode"
                        select
                        label="Индекс"
                        fullWidth
                        value={props.values.postcode}
                        onChange={props.onChange}
                    >
                        {
                            postcodes.map((postcode) => (
                                <MenuItem key={postcode} value={postcode}>{postcode}</MenuItem>
                            ))
                        }
                    </TextField>
                </Grid>
                <Grid item xs={4.5}>
                    <TextField
                        fullWidth
                        select
                        label="Должность"
                        id="position"
                        name="position"
                        value={props.values.position}
                        onChange={props.onChange}
                    >
                        <MenuItem value="">Не выбрано</MenuItem>
                        <MenuItem value={'Почтальон'}>Почтальон</MenuItem>
                        <MenuItem value={'Водитель'}>Водитель</MenuItem>
                        <MenuItem value={'Сотрудник отделения связи'}>Сотрудник отделения связи</MenuItem>
                        <MenuItem value={'Сотрудник сортировочного центра'}>Сотрудник сортировочного центра</MenuItem>
                    </TextField>
                </Grid>
                <Grid item xs={3}>
                    <LocalizationProvider dateAdapter={AdapterMoment}>
                        <DesktopDatePicker
                            name="date_start"
                            onChange={handleDateStartChange}
                            value={props.values.birthday_start}
                            label={"Дата рождения С"}
                            inputFormat="DD/MM/YYYY"
                            renderInput={(params) => <TextField fullWidth {...params} />}
                        />
                    </LocalizationProvider>
                </Grid>
                <Grid item xs={3}>
                    <LocalizationProvider dateAdapter={AdapterMoment}>
                        <DesktopDatePicker
                            name="date_end"
                            onChange={handleDateEndChange}
                            value={props.values.birthday_end}
                            label={"Дата рождения По"}
                            inputFormat="DD/MM/YYYY"
                            renderInput={(params) => <TextField fullWidth {...params} />}
                        />
                    </LocalizationProvider>
                </Grid>
                <Grid item xs={2}>
                    <TextField
                        fullWidth
                        select
                        label="Пол"
                        name="position"
                        value={props.values.sex}
                        onChange={props.onChange}
                    >
                        <MenuItem value="">Не выбрано</MenuItem>
                        <MenuItem value={'Мужской'}>Мужской</MenuItem>
                        <MenuItem value={'Женский'}>Женский</MenuItem>
                    </TextField>
                </Grid>
                <Grid item xs={4}>
                    <MuiPhoneNumber defaultCountry={'ru'} onlyCountries={['ru']} onChange={handleNumberPhoneChange} />
                </Grid>
            </Grid>
        </div>
    );
}