import * as React from 'react';
import {useParams } from 'react-router-dom';
import {useState} from "react";
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import {Grid} from "@mui/material";
import axios from 'axios';
import Head from "../components/head";

import CheckIcon from '@mui/icons-material/Check';

import "../styles/order-tracking.css"
import {useEffect} from "react";

function OrderTracking() {
    let {orderId} = useParams();
    const [dataTracking, setDataTracking] = useState({
        type: "",
        status: "",
        stages: []
    });

    useEffect(() => {
        axios.get("http://localhost:8080/api/v1/sending", {
            params: {
                order_id: orderId
            }
        })
        .then(
            (response) => {
                setDataTracking(response.data);
            }
        )
    }, [])

    function generateTrack(stages) {
        return stages.map((stage) => {
            return <ListItem key={stage.timestamp}>
                        <ListItemText
                            primary={stage.name}
                            secondary={stage.date + ", " + stage.postcode}
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
            return <p style={{color: '#6C6D6F'}}>{"Вручено " + stage.date}</p>;
        }
        return null;
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
                            <p style={{float: 'right'}}>{orderId}</p>
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
