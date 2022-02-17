import React, { useRef, useState } from 'react';
import LeagueTable from "./LeagueTable";
import MatchesTable from "./MatchesTable";
import ChampionshipPrediction from "./ChampionshipPrediction";
import { Button } from "react-bootstrap";
import { Row, Col } from 'react-bootstrap';

import axios from 'axios';


function LeagueInformationContainer(props) {
	const leagueComponentRef = useRef();
	const matchesTableRef = useRef();

	const playNextWeek = () => {
		axios.post("http://localhost:8080/matches").then(
			res => {
				console.log(res.status);
				console.log(res.data);
				leagueComponentRef.current.update();
			}
		)
	};

	const resetLeague = () => {
		axios.delete("http://localhost:8080/league-table").then(
			res => {
				console.log(res.status);
				console.log(res.data);
			}
		)
	};


	return (
		<div>
			<Row>
				<Col>
					<LeagueTable ref={leagueComponentRef} />
				</Col>
				<Col>
					<MatchesTable ref={matchesTableRef} />
				</Col>
				<Col>
					<ChampionshipPrediction />
				</Col>
			</Row>
			<Row>
				<Col>
					<Button variant="primary">Play All Weeks</Button>
				</Col>
				<Col>
					<Button variant="info" onClick={playNextWeek}>Play Next Week</Button>
				</Col>
				<Col>
					<Button variant="danger" onClick={resetLeague}>Reset</Button>
				</Col>
			</Row>
		</div>
	);
}

export default LeagueInformationContainer;
