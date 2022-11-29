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
import axios from "axios";
import {useNavigate} from "react-router-dom";

import AlertDialog from "../components/alertDialog";

function Main() {
    const [orderId, setOrderId] = useState("");
    const [dialogIsOpen, setDialogIsOpen] = useState(false);
    const [dialogText, setDialogText] = useState("");

    const openDialog = () => setDialogIsOpen(true);
    const closeDialog = () => setDialogIsOpen(false);

    const navigate = useNavigate();

    const handleOrderIdChange = (e) => {
        setOrderId(e.target.value)
    }

    const clickSearchButton = () => {
        axios.get("http://localhost:8080/api/v1/sending", {
            params: {
                order_id: orderId
            }
        })
        .then(
            (response) => {
                navigate('/orders/' + orderId, {replace: false})
            })
        .catch(
            (error) => {
                setDialogText('Неверное значение order-id: ' + orderId);
                openDialog();
            }
        )
    }

    return (
        <div className="MainPage">
            <div className="headMain">
                <Grid container>
                    <Grid xs={2}>
                        <img src={logo} className="logo"/>
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
                    onClick={clickSearchButton}
                >
                </Button>
            </Stack>
            <Slide direction="right" in={true} timeout={1500}>
                <img className="background" src={background}/>
            </Slide>
            <AlertDialog open={dialogIsOpen} onClose={closeDialog} text={dialogText}/>
        </div>
    );
}

export default Main;
