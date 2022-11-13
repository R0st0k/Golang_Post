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

export default function Registration(){
    const [type, setType] = React.useState("");

    const handleChange = (event) => {
        setType(event.target.value);
    };

    return (
        <>
            <Head />
            <Card>
                <CardContent>
                    <p>Регистрация нового отправления</p>

                    <Stack direction="row" spacing={1}>
                        <Typography>
                            1. Выберите тип посылки
                        </Typography>
                        <Box sx={{
                            width: 150
                        }}>
                        <TextField
                            fullWidth
                            select
                            label="Тип"
                            id="type"
                            value={type}
                            label="Тип"
                            onChange={handleChange}
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
                                label="Фамилия"
                                fullWidth
                                />
                        </Grid>
                        <Grid item xs={3}>
                            <TextField
                                label="Имя"
                                fullWidth
                            />
                        </Grid>
                        <Grid item xs={3}>
                            <TextField
                                label="Отчетсво"
                                fullWidth
                            />
                        </Grid>
                        <Grid item xs={3}></Grid>
                        <Grid item xs={3}>
                            <TextField
                                label="Город"
                                fullWidth
                            />
                        </Grid>
                        <Grid item xs={6}>
                            <TextField
                                label="Улица/дом/корпус/квартира"
                                fullWidth
                            />
                        </Grid>
                        <Grid item xs={3}>
                            <TextField
                                label="Индекс"
                                fullWidth
                            />
                        </Grid>
                    </Grid>
                </CardContent>
            </Card>
        </>
    )
}