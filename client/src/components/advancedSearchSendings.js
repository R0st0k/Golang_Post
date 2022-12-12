import {FormControl, Grid, Typography} from "@mui/material";
import TextField from "@mui/material/TextField";
import MenuItem from "@mui/material/MenuItem";
import Select from '@mui/material/Select';
import Chip from '@mui/material/Chip';
import Slider from '@mui/material/Slider';
import InputLabel from '@mui/material/InputLabel';
import Stack from "@mui/material/Stack";
import Box from "@mui/material/Box";
import CloseIcon from "@mui/icons-material/Close";
import { AdapterMoment } from '@mui/x-date-pickers/AdapterMoment';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DesktopDatePicker } from '@mui/x-date-pickers/DesktopDatePicker';
import * as React from "react";
import {useEffect, useState} from "react";
import axios from "axios";

export default function AdvancedSearchSendings(props){
    const [cityAndPostcodes, setCityAndPostcodes] = useState([]);

    useEffect(() => {
        axios.get('http://localhost:8080/api/v1/postcodes_by_settlement', {
            params: {
                type: "Отделение связи"
            }
        })
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

    const handleDateStartChange = value =>{
        let event = {};
        const target = {name: "date_start", value: value.format('YYYY-M-DD')};
        event.target = target;
        props.onChange(event);
    }

    const handleDateEndChange = value =>{
        let event = {};
        const target = {name: "date_finish", value: value.format('YYYY-M-DD')};
        event.target = target;
        props.onChange(event);
    }

    return (
        <div>
            <Grid container spacing={3}>
                <Grid item xs={3}>
                    <FormControl fullWidth>
                    <InputLabel id="type-select">Тип</InputLabel>
                    <Select
                        labelId="type-select"
                        label="Тип"
                        multiple
                        fullWidth
                        name="type"
                        value={props.values.type}
                        onChange={props.onChange}
                        renderValue={(selected) => (
                            <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                                {selected.map((value) => (
                                    <Chip key={value} label={value} />
                                ))}
                            </Box>
                        )}
                    >
                        <MenuItem value={'Письмо'}>Письмо</MenuItem>
                        <MenuItem value={'Посылка'}>Посылка</MenuItem>
                        <MenuItem value={'Бандероль'}>Бандероль</MenuItem>
                    </Select>
                    </FormControl>
                </Grid>
                <Grid item xs={3}>
                    <LocalizationProvider dateAdapter={AdapterMoment}>
                        <DesktopDatePicker
                            name="date_start"
                            onChange={handleDateStartChange}
                            value={props.values.date_start}
                            label={"Дата регистрации С"}
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
                            value={props.values.date_finish}
                            label={"Дата регистрации По"}
                            inputFormat="DD/MM/YYYY"
                            renderInput={(params) => <TextField fullWidth {...params} />}
                        />
                    </LocalizationProvider>
                </Grid>
                <Grid item xs={3}>
                    <FormControl fullWidth>
                        <InputLabel id="status-select">Статус</InputLabel>
                        <Select
                            labelId="status-select"
                            multiple
                            fullWidth
                            name="status"
                            label="Статус"
                            value={props.values.status}
                            onChange={props.onChange}
                            renderValue={(selected) => (
                                <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                                    {selected.map((value) => (
                                        <Chip key={value} label={value} />
                                    ))}
                                </Box>
                            )}
                        >
                            <MenuItem value={'В пути'}>В пути</MenuItem>
                            <MenuItem value={'Доставлено'}>Доставлено</MenuItem>
                            <MenuItem value={'Утеряно'}>Утеряно</MenuItem>
                        </Select>
                    </FormControl>
                </Grid>
                <Grid item xs={4.5}>
                    <TextField
                        fullWidth
                        name={"sender_settlement"}
                        label={"Откуда"}
                        value={props.values.sender_settlement}
                        onChange={props.onChange}
                    />
                    {/*<Autocomplete
                        renderInput={(params => <TextField
                            {...params}
                            label={"Откуда"}
                        />)}
                        onChange={handleChangeCity("sender_settlement")}
                        options={Object.keys(cityAndPostcodes)}
                    />*/}
                </Grid>
                <Grid item xs={4.5}>
                    <TextField
                        fullWidth
                        name={"receiver_settlement"}
                        label={"Куда"}
                        value={props.values.receiver_settlement}
                        onChange={props.onChange}
                    />
                    {/*<Autocomplete
                        renderInput={(params => <TextField
                            {...params}
                            label={"Куда"}
                        />)}
                        onChange={handleChangeCity("receiver_settlement")}
                        options={Object.keys(cityAndPostcodes)}
                    />*/}
                </Grid>
                <Grid item xs={3}>
                        <Typography textAlign={"center"} gutterBottom>
                            Вес
                        </Typography>
                        <Slider
                            value={props.values.weight}
                            onChange={props.onChange}
                            name="weight"
                            min={10}
                            max={1000}
                            step={10}
                            valueLabelDisplay="auto"
                        />
                </Grid>
                <Grid item xs={4}>
                    <Typography textAlign={"center"} gutterBottom>
                        Длина
                    </Typography>
                    <Slider
                        value={props.values.length}
                        onChange={props.onChange}
                        name="length"
                        min={10}
                        max={1000}
                        step={10}
                        valueLabelDisplay="auto"
                    />
                </Grid>
                <Grid item xs={4}>
                    <Typography textAlign={"center"} gutterBottom>
                        Ширина
                    </Typography>
                    <Slider
                        value={props.values.width}
                        onChange={props.onChange}
                        name="width"
                        min={10}
                        max={1000}
                        step={10}
                        valueLabelDisplay="auto"
                    />
                </Grid>
                <Grid item xs={4}>
                    <Typography textAlign={"center"} gutterBottom>
                        Высота
                    </Typography>
                    <Slider
                        value={props.values.height}
                        onChange={props.onChange}
                        name="height"
                        min={1}
                        max={1000}
                        step={10}
                        valueLabelDisplay="auto"
                    />
                </Grid>
            </Grid>
        </div>
    );
}