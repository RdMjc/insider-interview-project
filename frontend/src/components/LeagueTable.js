import React, { useEffect, useState } from 'react';
import { Table } from 'react-bootstrap';
import axios from 'axios';

function LeagueTable(props) {
	const [teams, setTeams] = useState([]);

	const getLeagueTable = () => {
		axios.get("http://localhost:8080/league-table").then(
			res => {
				setTeams(res.data["teams"]);
			}
		)
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
				{
					teams.map((team, index) => {
						// console.log(team);
						// console.log(Object.keys(team));
						return (
							<tr>
								<td>{Object.keys(team)}</td>
								<td>{team[Object.keys(team)[0]]["points"]}</td>
								<td>{team[Object.keys(team)[0]]["games"]}</td>
								<td>{team[Object.keys(team)[0]]["wins"]}</td>
								<td>{team[Object.keys(team)[0]]["draws"]}</td>
								<td>{team[Object.keys(team)[0]]["losses"]}</td>
								<td>{team[Object.keys(team)[0]]["goalDifference"]}</td>
							</tr>
						);
					})
				}

			</tbody>
		</Table>
	);
}

export default LeagueTable;
