import * as React from 'react';
import { Routes, Route, useParams } from 'react-router-dom';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import {Grid} from "@mui/material";

import Head from "../components/head";

import CheckIcon from '@mui/icons-material/Check';

import "../styles/order-tracking.css"

function OrderTracking() {
    let {orderId} = useParams();

    const dataTracking = {
        order_id: "123e4567-e89b-12d3-a456-426655440000",
        type: "Письмо",
        status: "Доставлено",
        stages: [
            {
                name: "Принято в отделении",
                postcode: "123456",
                timestamp: "2012-07-14T01:00:00+01:00"
            },
            {
                name: "Покинуло место приёма",
                postcode: "234531",
                timestamp: "2014-07-14T01:00:00+01:00"
            },
            {
                name: "Прибыло в сортировочный центр",
                postcode: "236445",
                timestamp: "2013-07-14T01:00:00+02:00"
            },
        ]
    }

    function generateTrack(stages) {
        return stages.map((stage) => {
            return <ListItem key={stage.timestampe}>
                        <ListItemText
                            primary={stage.name}
                            secondary={dateFormat(stage.timestamp) + ", " + stage.postcode}
                        />
                    </ListItem>
        });
    }

    function trackStatus(status){
        if(status === "Доставлено"){
            return <CheckIcon style={{float: 'right'}} color='success' />;
        }
        return null;
    }

    function trackIsComplete(status, stage){
        if(status === "Доставлено"){
            return <p style={{color: '#6C6D6F'}}>{"Вручено " + dateFormat(stage.timestamp)}</p>;
        }
        return null;
    }

    function dateFormat(dateString){
        const date = new Date(dateString);
        const options = {day: 'numeric', month: 'long', year: 'numeric'};
        return date.toLocaleDateString('ru-RU', options) + ", "+ date.toLocaleTimeString('ru-RU');

    }

    return (
        <div>
            <Head />
            <div className="Track">
                <div className="TrackHead">
                    <Grid container spacing={0}>
                        <Grid item xs={8}>
                            <p style={{color: '#2971C5'}}>{dataTracking.type}</p>
                        </Grid>
                        <Grid item xs={4}>
                            {trackStatus(dataTracking.status)}
                        </Grid>
                        <Grid item xs={4}>
                            {trackIsComplete(dataTracking.status, dataTracking.stages[0])}
                        </Grid>
                        <Grid item xs={8}>
                            <p style={{float: 'right'}}>{dataTracking.order_id}</p>
                        </Grid>
                    </Grid>
                </div>
                <List dense={1}>
                    {generateTrack(dataTracking.stages)}
                </List>
            </div>
        </div>
    );
}

export default OrderTracking;
