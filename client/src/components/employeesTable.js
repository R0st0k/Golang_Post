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
import {useNavigate} from "react-router-dom";

const headCells = [
    {
        id: 'name',
        numeric: false,
        disableSort: false,
        label: 'ФИО',
    },
    {
        id: 'sex',
        numeric: false,
        disableSort: false,
        label: 'Пол',
    },
    {
        id: 'birthday',
        numeric: false,
        disableSort: false,
        label: 'Дата рождения',
    },
    {
        id: 'position',
        numeric: false,
        disableSort: false,
        label: 'Должность',
    },
    {
        id: 'phone_number',
        numeric: false,
        disableSort: true,
        label: 'Номер телефона',
    },
    {
        id: 'settlement',
        numeric: true,
        disableSort: false,
        label: 'Город',
    },
    {
        id: 'postcode',
        numeric: false,
        disableSort: false,
        label: 'Индекс отделения',
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
                        {headCell.disableSort ? headCell.label :
                            <TableSortLabel
                                active={orderBy === headCell.id}
                                direction={orderBy === headCell.id ? order : 'asc'}
                                onClick={createSortHandler(headCell.id)}
                            >
                                {headCell.label}

                            </TableSortLabel>
                        }
                    </TableCell>
                ))}
            </TableRow>
        </TableHead>
    );
}


export default function EmployeesTable(props) {

    const [selected, setSelected] = React.useState([]);

    const order = props.order;
    const orderBy = props.orderBy;
    const page = props.page;
    const rowsPerPage = props.rowsPerPage;
    const rows = props.data;
    const total = props.total;

    const isSelected = (name) => selected.indexOf(name) !== -1;

    const handleChangePage = (event, value) => {
        const target = {name: "page", value: value};
        props.handleChangeTable(target);
    }

    const handleChangeRowsPerPage = (event) => {
        props.onChangeRowsPerPage(event.target.value);
    }

    const handleRequestSort = (event, property) => {
        const isAsc = orderBy === property && order === 'asc';
        const target_1 = {name:  'order', value: isAsc ? 'desc' : 'asc'};
        const target_2 = {name:  'orderBy', value: property};
        props.handleChangeTable(target_1);
        props.handleChangeTable(target_2);
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
                                            aria-checked={isItemSelected}
                                            tabIndex={-1}
                                            key={row.order_id}
                                            selected={isItemSelected}
                                        >
                                            <TableCell
                                                component="th"
                                                scope="row"
                                                padding="normal"
                                                align="center"
                                            >
                                                {row.surname} {row.name} {row.middle_name}
                                            </TableCell>
                                            <TableCell align="center">{row.sex}</TableCell>
                                            <TableCell align="center">{row.birthday}</TableCell>
                                            <TableCell align="center">{row.position}</TableCell>
                                            <TableCell align="center">{row.phone_number}</TableCell>
                                            <TableCell align="center">{row.settlement}</TableCell>
                                            <TableCell align="center">{row.postcode}</TableCell>
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
                    count={total}
                    rowsPerPage={rowsPerPage}
                    page={page}
                    onPageChange={handleChangePage}
                    onRowsPerPageChange={handleChangeRowsPerPage}
                />
            </Paper>
        </Box>
    );
}