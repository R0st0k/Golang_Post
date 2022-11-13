import React, {useCallback} from 'react';
import {useNavigate} from 'react-router-dom';
import Button from '@mui/material/Button';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';
import Fade from '@mui/material/Fade';
import KeyboardArrowDownIcon from '@mui/icons-material/KeyboardArrowDown';

import "../styles/fade-menu.css"

export default function FadeMenu() {
    const [anchorEl, setAnchorEl] = React.useState(null);
    const navigate = useNavigate();
    const open = Boolean(anchorEl);
    const handleClick = (event) => {
        setAnchorEl(event.currentTarget);
    };
    const handleClose = () => {
        setAnchorEl(null);
    };
    const handleRegistration = useCallback(() => navigate('/registration', {replace: true}), [navigate]);
    const handleStatics = useCallback(() => navigate('/', {replace: true}), [navigate]);
    const handleTable = useCallback(() => navigate('/', {replace: true}), [navigate]);

    return (
        <div>
            <Button
                id="fade-button"
                aria-controls={open ? 'fade-menu' : undefined}
                aria-haspopup="true"
                aria-expanded={open ? 'true' : undefined}
                onClick={handleClick}
                endIcon={<KeyboardArrowDownIcon />}
                className="fadeButton"
            >
                    Отправления
            </Button>
            <Menu
                id="fade-menu"
                MenuListProps={{
                    'aria-labelledby': 'fade-button',
                }}
                anchorEl={anchorEl}
                open={open}
                onClose={handleClose}
                TransitionComponent={Fade}
            >
                <MenuItem onClick={handleRegistration}>Новое отправление</MenuItem>
                <MenuItem onClick={handleStatics}>Статистика</MenuItem>
                <MenuItem onClick={handleTable}>Информация об отправлениях</MenuItem>
            </Menu>
        </div>
    );
}