import React, { useState,useEffect } from 'react';
import { withStyles, makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';

const useStyles = makeStyles({
  table: {
    minWidth: '100%',
  },
});

const StyledTableHeaderCell = withStyles((theme) => ({
    head: {
      backgroundColor: '#cff1ef',
      color: theme.palette.common.black,
      fontSize: 18,
      fontWeight: 'bold',
    },
    body: {
      fontSize: 14,
    },
  }))(TableCell);

  const StyledTableContentCell = withStyles((theme) => ({
    body: {
      fontSize: 16,
    },
  }))(TableCell);
  
function createData(transaction_date, bank, description, amount) {
  return { transaction_date, bank, description, amount };
}

export default function SimpleTable(props) {
  const classes = useStyles();
  const [rows , setRows] = useState([]);


  useEffect(()=>{
    var rows = [] 
    function createRows(){
        if (typeof props.data !== undefined) {

            for (var v in props.data){
                rows.push(createData(props.data[v].transaction_date, props.data[v].bank, props.data[v].description , props.data[v].amount))
            }
        }

        setRows(rows)

    }   
    createRows()
  },[props.data]);

  return (
    <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="a dense table">
        <TableHead>
          <TableRow>
            <StyledTableHeaderCell>Transaction Date</StyledTableHeaderCell>
            <StyledTableHeaderCell align="right">Bank</StyledTableHeaderCell>
            <StyledTableHeaderCell align="right">Description</StyledTableHeaderCell>
            <StyledTableHeaderCell align="right">Amount</StyledTableHeaderCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {rows.map((row) => (
            <TableRow key={row.name}>
              <StyledTableContentCell component="th" scope="row">
                {row.transaction_date}
              </StyledTableContentCell>
              <StyledTableContentCell align="right">{row.bank}</StyledTableContentCell>
              <StyledTableContentCell align="right">{row.description}</StyledTableContentCell>
              <StyledTableContentCell align="right">{row.amount}</StyledTableContentCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
