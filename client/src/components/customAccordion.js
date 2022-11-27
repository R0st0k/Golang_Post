import { styled } from '@mui/material/styles';
import ArrowForwardIosSharpIcon from '@mui/icons-material/ArrowForwardIosSharp';
import MuiAccordion from '@mui/material/Accordion';
import MuiAccordionSummary from '@mui/material/AccordionSummary';
import MuiAccordionDetails from '@mui/material/AccordionDetails';
import * as React from "react";
import {Typography} from "@mui/material";

const Accordion = styled((props) => (
    <MuiAccordion disableGutters elevation={0} square {...props} />
))(() => ({

    '&:not(:last-child)': {
        borderBottom: 0,
    },
    '&:before': {
        display: 'none',
    },
}));

const AccordionSummary = styled((props) => (
    <MuiAccordionSummary
        expandIcon={<ArrowForwardIosSharpIcon sx={{ fontSize: '0.9rem' }} />}
        {...props}
    />
))(({ theme }) => ({
    flexDirection: 'row-reverse',
    '& .MuiAccordionSummary-expandIconWrapper.Mui-expanded': {
        transform: 'rotate(90deg)',
    },
    '& .MuiAccordionSummary-content': {
        marginLeft: theme.spacing(1),
    },
}));

const AccordionDetails = styled(MuiAccordionDetails)(({ theme }) => ({
    padding: theme.spacing(2),
    borderTop: '1px solid rgba(0, 0, 0, .125)',
}));

export default function CustomAccordion(props) {
    const [expanded, setExpanded] = React.useState(false);

    const handleChange = (event, newExpanded) => {
        setExpanded(!!newExpanded);
    };

    return (
        <div>
            <Accordion expanded={expanded} onChange={handleChange}>
                <AccordionSummary >
                    <Typography>{props.label}</Typography>
                </AccordionSummary>
                <AccordionDetails>
                    {props.element}
                </AccordionDetails>
            </Accordion>
        </div>
    );
}