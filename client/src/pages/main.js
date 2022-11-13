import React, { useState } from 'react';
import InputBase from '@mui/material/InputBase';
import Button from '@mui/material/Button';
import Stack from '@mui/material/Stack';
import SearchIcon from '@mui/icons-material/Search';
import FadeMenu from "../components/fade-menu";
import "../styles/main.css"
import Grid from "@mui/material/Unstable_Grid2";
import Box from '@mui/material/Box';

import background from '../images/smile_harold.jpg';
import logo from '../images/logo.jpg'
import Slide from '@mui/material/Slide';

function Main() {
    const [orderId, setOrderId] = useState("");

    const handleOrderIdChange = (e) => {
        setOrderId(e.target.value)
    }

    return (
        <div className="MainPage">
            <div className="headMain">
                <Grid container>
                    <Grid xs={2}>
                        <img src={logo} className="logo"></img>
                    </Grid>
                    <Grid xs={6}></Grid>
                    <Grid xs={2}>
                        <Box mt={2}>
                        <Button
                            href={'/'}
                            className="fadeLink"
                        >
                            Сотрудники
                        </Button>
                        </Box>
                    </Grid>
                    <Grid xs={2}>
                        <Box mt={2}>
                        <FadeMenu />
                        </Box>
                    </Grid>
                </Grid>
            </div>
            <Stack spacing={1} direction="row" className="InputSearch">
                <InputBase
                    className="OrderSearch"
                    placeholder="Введите значение order-id"
                    value={orderId}
                    onChange={handleOrderIdChange}
                />
                <Button
                    variant="contained"
                    className="IconSearch"
                    startIcon={<SearchIcon />}
                    href={'orders/' + orderId}
                >
                </Button>
            </Stack>
            <Slide direction="right" in={true} timeout={1500}>
                <img className="background" src={background}></img>
            </Slide>

        </div>
    );
}

export default Main;
