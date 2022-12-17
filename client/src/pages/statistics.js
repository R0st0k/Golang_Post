import * as React from 'react';
import Head from "../components/head";
import Autocomplete from '@mui/material/Autocomplete';
import {useEffect, useState} from "react";
import Stack from "@mui/material/Stack";
import Grid from '@mui/material/Unstable_Grid2';

import 'react-checkbox-tree/lib/react-checkbox-tree.css';
import MenuItem from "@mui/material/MenuItem";
import TextField from "@mui/material/TextField";
import axios from "axios";
import qs from 'qs';
import Button from "@mui/material/Button";
import Box from "@mui/material/Box";


import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    BarElement,
    Title,
    Tooltip,
    Legend,
} from 'chart.js';
import { Bar } from 'react-chartjs-2';

ChartJS.register(
    CategoryScale,
    LinearScale,
    BarElement,
    Title,
    Tooltip,
    Legend
);

const options = {
    responsive: true,
};

export default function Statistics(props){
    const [settlement, setSettlement] = useState([]);
    const [typeOfSending, setTypeOfSending] = useState([]);
    const [direction, setDirection] = useState("");
    const [typeOfStatistic, setTypeOfStatistic] = useState("");
    const [statisticsData, setStatisticData] = useState({
        labels: ['empty'],
        datasets: [
            {
                data: [0],
                backgroundColor: 'rgba(255, 99, 132, 0.5)',
            }
        ]
    });
    const [showBar, setShowBar] = useState(false);

    const [cityAndPostcodes, setCityAndPostcodes] = useState([]);

    useEffect(() => {
        axios.get('http://localhost:8080/api/v1/postcodes_by_settlement')
            .then(
                (response) => {
                    setCityAndPostcodes(response.data);
                }
            )
    }, []);

    const typeOptions = [ "Письмо", "Бандероль", "Посылка"];
    const dic = {
        "Количество": "Количество единиц",
        "Время":"Часы",
        "Вес":"Граммы",
    }

    const handleClickButtonShow = (event) => {
        event.preventDefault();
        const api = axios.create({
            paramsSerializer: {
                serialize: (params) => qs.stringify(params, {arrayFormat: 'repeat'})
            }
        });
        api.get('http://localhost:8080/api/v1/sending_statistics', {
            params: {
                settlement: settlement,
                type: typeOfSending,
                direction: direction,
                statistics: typeOfStatistic
            }
        })
            .then(
                (response) => {
                    const data = response.data;
                    const labels = data.map((element) => element.key);
                    const statistics = {
                        labels,
                        datasets: [
                            {
                                label: dic[typeOfStatistic],
                                data: data.map((element) => element.value),
                                backgroundColor: 'rgba(49,78,253,0.5)',
                            }
                        ]
                    };
                    setStatisticData(statistics);
                    setShowBar(true);
                }
            )
    }

    return (
        <>
            <Head/>
            <Box ml={5} mt={2}>
                <form
                    onSubmit={handleClickButtonShow}
                    autoComplete="off"
                >
                <Grid spacing={10} container>
                    <Grid xs={3}>
                        <Stack spacing={2}>
                            <Autocomplete
                                multiple
                                renderInput={(params => <TextField
                                    {...params}
                                    label={"Населенный пункт"}
                                    required={settlement.length === 0}
                                />)}
                                onChange={(e, value) => setSettlement(value)}
                                options={Object.keys(cityAndPostcodes)}
                                freeSolo={settlement.length >= 10 ? false : true}
                                getOptionDisabled={(options) => (settlement.length >= 10 ? true : false)}
                            />

                            <Autocomplete
                                multiple
                                renderInput={(params => <TextField
                                    {...params}
                                    label={"Тип"}
                                    required={typeOfSending.length === 0}
                                />)}
                                onChange={(e, value) => setTypeOfSending(value)}
                                options={typeOptions}
                            />
                            <TextField
                                required
                                select
                                label="Направление"
                                name="direction"
                                value={direction}
                                onChange={(e) => setDirection(e.target.value)}
                            >
                                <MenuItem value={'Отправления'}>Отправление</MenuItem>
                                <MenuItem value={'Получения'}>Получение</MenuItem>
                            </TextField>
                            <TextField
                                required
                                select
                                label="Данные"
                                name="type_of_statistic"
                                value={typeOfStatistic}
                                onChange={(e) => setTypeOfStatistic(e.target.value)}
                            >
                                <MenuItem value={'Количество'}>Количество</MenuItem>
                                <MenuItem value={'Время'}>Среднее время в пути</MenuItem>
                                <MenuItem value={'Вес'}>Средний вес</MenuItem>
                            </TextField>
                        </Stack>
                    </Grid>
                    <Grid xs={6}>
                        <Button variant={"contained"} type={"submit"}>
                            Показать
                        </Button>
                        {
                            showBar ? <Bar options={options} data={statisticsData} /> : null
                        }
                    </Grid>
                </Grid>
                </form>
            </Box>
        </>
    )
}