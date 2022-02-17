import React, {useEffect, useState} from 'react';
import { Table } from 'react-bootstrap';
import axios from 'axios';

function LeagueTable(props) {

	const getLeagueTable = () => {

	};

	// takes place everytime this component is rendered
	useEffect(() => {
		getLeagueTable()
	}, []);
	
	return (
		<Table responsive="sm">
			<thead>
				<tr>
					<th>Teams</th>
					<th>Points</th>
					<th>Games</th>
					<th>Wins</th>
					<th>Draws</th>
					<th>Losses</th>
					<th>GD</th>
				</tr>
			</thead>
			<tbody>
				<tr>
					<td>1</td>
					<td>Table cell</td>
					<td>Table cell</td>
					<td>Table cell</td>
					<td>Table cell</td>
					<td>Table cell</td>
					<td>Table cell</td>
				</tr>
			</tbody>
		</Table>
	);
}

export default LeagueTable;
