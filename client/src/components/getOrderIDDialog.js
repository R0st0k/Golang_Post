import * as React from 'react';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import Button from '@mui/material/Button';

export default function GetOrderIDDialog(props) {
    const { onClose, open } = props;

    const orderId = 123123;

    return (
        <Dialog open={open} onClose={onClose}>
            <DialogContent>
                <DialogContentText textAlign={'center'}>
                    Отправление было успешно зарегистрировано.
                    <br/>
                    order-id: {orderId}
                </DialogContentText>
            </DialogContent>
            <DialogActions>
                <Button variant="outlined" href={'/registration'}>Зарегистрировать новое отправление</Button>
                <Button autoFocus variant="contained" href={'/'}>
                    На главную страницу
                </Button>
            </DialogActions>
        </Dialog>
    );
}
