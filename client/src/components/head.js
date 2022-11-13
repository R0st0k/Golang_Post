import * as React from 'react';
import Grid from '@mui/material/Unstable_Grid2';
import Button from '@mui/material/Button';
import FadeMenu from "./fade-menu";

import "../styles/head.css"

export default function Head() {

    return (
        <div>
            <Grid container>
                <Grid xs={8}>
                    <Button
                        href={'/'}
                        className="fadeLink"
                    >
                        На главную страницу
                    </Button>
                </Grid>
                <Grid xs={2}>
                    <Button
                        href={'/'}
                        className="fadeLink"
                    >
                        Сотрудники
                    </Button>
                </Grid>
                <Grid xs={2}>
                    <FadeMenu />
                </Grid>
            </Grid>
        </div>
    )
}