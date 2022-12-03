import * as React from 'react';
import Head from "../components/head";
import Autocomplete from '@mui/material/Autocomplete';
import CheckboxTree from 'react-checkbox-tree';
import {useEffect, useState} from "react";
import Chip from '@mui/material/Chip';
import Stack from "@mui/material/Stack";
import Grid from '@mui/material/Unstable_Grid2';

import 'react-checkbox-tree/lib/react-checkbox-tree.css';
import MenuItem from "@mui/material/MenuItem";
import TextField from "@mui/material/TextField";
import axios from "axios";
import {ButtonGroup} from "@mui/material";
import Button from "@mui/material/Button";
import Box from "@mui/material/Box";

export default function Statistics(props){
    const [settlement, setSettlement] = useState([]);
    const [typeOfSending, setTypeOfSending] = useState([]);
    const [expandedTypeOfSending, setExpandedTypeOfSending] = useState([]);
    const [direction, setDirection] = useState("");
    const [typeOfStatistic, setTypeOfStatistic] = useState("");

    const [cityAndPostcodes, setCityAndPostcodes] = useState([]);

    useEffect(() => {
        axios.get('http://localhost:8080/api/v1/postcodes_by_settlement')
            .then(
                (response) => {
                    setCityAndPostcodes(response.data);
                }
            )
    }, []);

    const nodesTypesOfSending = [{
        value: 'types_of_sendings',
        label: 'Тип',
        children: [
            { value: 'Письмо', label: 'Письмо' },
            { value: 'Бандероль', label: 'Бандероль' },
            { value: 'Посылка', label: 'Посылка' }
        ],
    }];

    const handleCheckTypesOfSending = (value) => {
        setTypeOfSending(value);
    }

    const handleExpandTypesOfSending = (value) => {
        setExpandedTypeOfSending(value);
    }

    const handleClickButtonShow = (event) => {

    }

    return (
        <>
            <Head/>
            <Box ml={5} mt={2}>
                <Grid spacing={10} container>
                    <Grid xs={3}>
                        <Stack spacing={2}>
                            <Autocomplete
                                multiple
                                renderInput={(params => <TextField
                                    {...params}
                                    label={"Населенный пункт"}
                                />)}
                                onChange={(e, value) => setSettlement(value)}
                                options={Object.keys(cityAndPostcodes)}
                            />
                            <CheckboxTree
                                nodes={nodesTypesOfSending}
                                checked={typeOfSending}
                                expanded={expandedTypeOfSending}
                                onCheck={handleCheckTypesOfSending}
                                onExpand={handleExpandTypesOfSending}
                                showNodeIcon={false}
                            />
                            <TextField

                                select
                                label="Направление"
                                name="direction"
                                value={direction}
                                onChange={(e) => setDirection(e.target.value)}
                            >
                                <MenuItem value={'Отправление'}>Отправление</MenuItem>
                                <MenuItem value={'Получение'}>Получение</MenuItem>
                            </TextField>
                            <TextField
                                select
                                label="Данные"
                                name="type_of_statistic"
                                value={typeOfStatistic}
                                onChange={(e) => setTypeOfStatistic(e.target.value)}
                            >
                                <MenuItem value={'Количество'}>Количество</MenuItem>
                                <MenuItem value={'Среднее время в пути'}>Среднее время в пути</MenuItem>
                                <MenuItem value={'Средний вес'}>Средний вес</MenuItem>
                            </TextField>
                        </Stack>
                    </Grid>
                    <Grid xs={6}>
                        <Button variant={"contained"} onClick={handleClickButtonShow}>
                            Показать
                        </Button>
                    </Grid>
                </Grid>
            </Box>
        </>
    )
}