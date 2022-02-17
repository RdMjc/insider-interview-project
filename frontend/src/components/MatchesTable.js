import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Table } from 'react-bootstrap';

function MatchesTable(props) {
	const [matches, setMatches] = useState([]);

	const getMatches = () => {
		axios.get("http://localhost:8080/matches").then(
			res => {
				setMatches(res.data["matches"]);
			}
		)
	};

	// takes place everytime this component is rendered
	useEffect(() => {
		getMatches()
	}, [props.refresh]);

	return (

		<Table striped hover>
			<thead>
				<tr>
					<th>Home Team</th>
					<th></th>
					<th></th>
					<th></th>
					<th>Away Team</th>
				</tr>
			</thead>
			<tbody>
				{
					matches.map((match, index) => {
						if (match.isPlayed) {
							return (
								<tr>
									<td>{match.homeTeam.name}</td>
									<td>{match.homeGoals}</td>
									<td>-</td>
									<td>{match.awayGoals}</td>
									<td>{match.awayTeam.name}</td>
								</tr>
							)
						} else {
							return (
								<tr>
									<td>{match.homeTeam.name}</td>
									<td></td>
									<td>vs</td>
									<td></td>
									<td>{match.awayTeam.name}</td>
								</tr>
							)
						}
					})
				}
			</tbody>
		</Table>
	);
}

export default MatchesTable;