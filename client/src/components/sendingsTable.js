import * as React from 'react';
import Box from '@mui/material/Box';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableHead from '@mui/material/TableHead';
import TableContainer from '@mui/material/TableContainer';
import TablePagination from '@mui/material/TablePagination';
import TableRow from '@mui/material/TableRow';
import TableSortLabel from '@mui/material/TableSortLabel';
import {Paper} from "@mui/material";
import {Typography} from "@mui/material";

function createData(orderID,type,date,departure_city, arrival_city,weight, shape,status) {
    return {
        orderID,
        type,
        date,
        departure_city,
        arrival_city,
        weight,
        shape,
        status
    };
}

const rows = [
    createData('1203884', 'Посылка', "5.10.22", "Санкт-Петербург", "Москва", 6000, "500x200x180", "В пути"),
    createData('2334745', 'Бандероль', "1.10.22", "Санкт-Петербург", "Оренбург", 300, "500x20x180", "Доставлено"),
];

const headCells = [
    {
        id: 'orderID',
        numeric: false,
        disablePadding: false,
        label: 'order-id',
    },
    {
        id: 'type',
        numeric: false,
        disablePadding: false,
        label: 'Тип',
    },
    {
        id: 'date',
        numeric: false,
        disablePadding: false,
        label: 'Дата регистрации отправления',
    },
    {
        id: 'departure_city',
        numeric: false,
        disablePadding: false,
        label: 'Откуда',
    },
    {
        id: 'arrival_city',
        numeric: false,
        disablePadding: false,
        label: 'Куда',
    },
    {
        id: 'weight',
        numeric: true,
        disablePadding: false,
        label: 'Вес',
    },
    {
        id: 'shape',
        numeric: false,
        disablePadding: false,
        label: 'Габариты',
    },
    {
        id: 'status',
        numeric: false,
        disablePadding: false,
        label: 'Статус',
    }
];

function EnhancedTableHead(props) {
    const { order, orderBy, onRequestSort } =
        props;
    const createSortHandler = (property) => (event) => {
        onRequestSort(event, property);
    };

    return (
        <TableHead>
            <TableRow>
                {headCells.map((headCell) => (
                    <TableCell
                        key={headCell.id}
                        align="center"
                        padding='normal'
                        sortDirection={orderBy === headCell.id ? order : false}
                    >
                        <TableSortLabel
                            active={orderBy === headCell.id}
                            direction={orderBy === headCell.id ? order : 'asc'}
                            onClick={createSortHandler(headCell.id)}
                        >
                            {headCell.label}

                        </TableSortLabel>
                    </TableCell>
                ))}
            </TableRow>
        </TableHead>
    );
}


export default function SendingsTable() {
    const [order, setOrder] = React.useState('asc');
    const [orderBy, setOrderBy] = React.useState('calories');
    const [selected, setSelected] = React.useState([]);
    const [page, setPage] = React.useState(0);
    const [rowsPerPage, setRowsPerPage] = React.useState(5);

    const isSelected = (name) => selected.indexOf(name) !== -1;

    const handleChangePage = () => {
        console.log('a')
    }

    const handleChangeRowsPerPage = () => {
        console.log('b')
    }

    const handleClick = (event, name) => {
        console.log('c')
    }

    const handleRequestSort = (event, property) => {
        const isAsc = orderBy === property && order === 'asc';
        setOrder(isAsc ? 'desc' : 'asc');
        setOrderBy(property);
    }

    return (
        <Box ml={"10%"} mt={5} sx={{ width: '80%' }}>
            <Paper sx={{ width: '100%', mb: 2 }}>
                <TableContainer>
                    <Table
                        sx={{ minWidth: 500 }}
                        size='medium'
                    >
                        <EnhancedTableHead
                            order={order}
                            orderBy={orderBy}
                            onRequestSort={handleRequestSort}
                        />
                        <TableBody>{
                            rows.slice(0, rowsPerPage)
                                .map((row) => {
                                const isItemSelected = isSelected(row.name);

                                return (
                                    <TableRow
                                        hover
                                        onClick={(event) => handleClick(event, row.name)}
                                        role="checkbox"
                                        aria-checked={isItemSelected}
                                        tabIndex={-1}
                                        key={row.orderID}
                                        selected={isItemSelected}
                                    >
                                        <TableCell
                                            component="th"
                                            scope="row"
                                            padding="normal"
                                            align="center"
                                        >
                                            {row.orderID}
                                        </TableCell>
                                        <TableCell align="center">{row.type}</TableCell>
                                        <TableCell align="center">{row.date}</TableCell>
                                        <TableCell align="center">{row.departure_city}</TableCell>
                                        <TableCell align="center">{row.arrival_city}</TableCell>
                                        <TableCell align="center">{row.weight}</TableCell>
                                        <TableCell align="center">{row.shape}</TableCell>
                                        <TableCell align="center">
                                            {
                                                row.status === "В пути" ? <Typography color="#ED8000">{row.status}</Typography> :  <Typography color="#6FD600">{row.status}</Typography>
                                            }
                                        </TableCell>
                                    </TableRow>
                                );
                            })
                        }
                        </TableBody>
                    </Table>
                </TableContainer>
                <TablePagination
                    rowsPerPageOptions={[5, 10, 25]}
                    component="div"
                    count={rows.length}
                    rowsPerPage={rowsPerPage}
                    page={page}
                    onPageChange={handleChangePage}
                    onRowsPerPageChange={handleChangeRowsPerPage}
                />
            </Paper>
        </Box>
    );
}