import {Grid, Typography} from "@mui/material";
import TextField from "@mui/material/TextField";
import MenuItem from "@mui/material/MenuItem";
import Stack from "@mui/material/Stack";
import Box from "@mui/material/Box";
import CloseIcon from "@mui/icons-material/Close";
import * as React from "react";

export default function AdvancedSearchSendings(props){

    const citysAndPostcodes =  [
        {city: "Москва", postcodes: [345123, 125234, 677521]},
        {city: "Питер",  postcodes: [235232, 235619]}
    ]

    return (
        <div>
            <Grid container spacing={2}>
                <Grid item xs={3}>
                    <TextField
                        fullWidth
                        select
                        label="Тип"
                        name="type"
                        value={props.values.type}
                        onChange={props.onChange}
                    >
                        <MenuItem value={'Письмо'}>Письмо</MenuItem>
                        <MenuItem value={'Посылка'}>Посылка</MenuItem>
                        <MenuItem value={'Бандероль'}>Бандероль</MenuItem>
                    </TextField>
                </Grid>
                <Grid item xs={6}>
                    <TextField
                        fullWidth
                        name="date"
                        label="Дата регистрации отправления"
                        value={props.values.date}
                        onChange={props.onChange}
                    />
                </Grid>
                <Grid item xs={3}>
                    <TextField
                        fullWidth
                        select
                        name="status"
                        label="Статус"
                        value={props.values.status}
                        onChange={props.onChange}
                    >
                        <MenuItem value={'inProgress'}>В пути</MenuItem>
                        <MenuItem value={'delivered'}>Доставлено</MenuItem>
                    </TextField>
                </Grid>
                <Grid item xs={4.5}>
                    <TextField
                        select
                        fullWidth
                        name="departure_city"
                        label="Откуда"
                        value={props.values.departure_city}
                        onChange={props.onChange}
                    >
                        {
                            citysAndPostcodes.map((city) =>{
                                return <MenuItem value={city.city}>{city.city}</MenuItem>
                            })
                        }
                    </TextField>
                </Grid>
                <Grid item xs={4.5}>
                    <TextField
                        select
                        fullWidth
                        name="arrival_city"
                        label="Куда"
                        value={props.values.arrival_city}
                        onChange={props.onChange}
                    >
                        {
                            citysAndPostcodes.map((city) =>{
                                return <MenuItem value={city.city}>{city.city}</MenuItem>
                            })
                        }
                    </TextField>
                </Grid>
                <Grid item xs={3}>
                    <Stack direction="row">
                        <TextField
                            inputProps={{ inputMode: 'numeric', pattern: '^[0-9]+(\\.[0-9]+)?$' }}
                            name="weight"
                            label="Вес"
                            fullWidth
                            value={props.values.weight}
                            onChange={props.onChange}
                        />
                        <Typography ml={1} mt={2} color={'#9F9B9B'}>
                            г
                        </Typography>
                    </Stack>
                </Grid>
                <Grid item xs={4}>
                    <Stack direction="row">
                        <TextField
                            inputProps={{ inputMode: 'numeric', pattern: '^[0-9]+(\\.[0-9]+)?$' }}
                            name="length"
                            label="Длина"
                            fullWidth
                            value={props.values.length}
                            onChange={props.onChange}
                        />
                        <Typography mt={2} color={'#9F9B9B'}>
                            мм
                        </Typography>
                        <Box mt={2}>
                            <CloseIcon style={{color: '#9F9B9B'}}/>
                        </Box>
                    </Stack>
                </Grid>
                <Grid item xs={4}>
                    <Stack direction="row">
                        <TextField
                            inputProps={{ inputMode: 'numeric', pattern: '^[0-9]+(\\.[0-9]+)?$' }}
                            name="width"
                            label="Ширина"
                            fullWidth
                            value={props.values.width}
                            onChange={props.onChange}
                        />
                        <Typography mt={2} color={'#9F9B9B'}>
                            мм
                        </Typography>
                        <Box mt={2}>
                            <CloseIcon style={{color: '#9F9B9B'}}/>
                        </Box>
                    </Stack>
                </Grid>
                <Grid item xs={4}>
                    <Stack direction="row">
                        <TextField
                            inputProps={{ inputMode: 'numeric', pattern: '^[0-9]+(\\.[0-9]+)?$' }}
                            name="height"
                            label="Высота"
                            fullWidth
                            value={props.values.height}
                            onChange={props.onChange}
                        />
                        <Typography mt={2} color={'#9F9B9B'}>
                            мм
                        </Typography>
                    </Stack>
                </Grid>
            </Grid>
        </div>
    );
}