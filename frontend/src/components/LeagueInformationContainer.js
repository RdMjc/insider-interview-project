import React, { useRef, useState } from 'react';
import LeagueTable from "./LeagueTable";
import MatchesTable from "./MatchesTable";
import ChampionshipPrediction from "./ChampionshipPrediction";
import { Button } from "react-bootstrap";
import { Row, Col } from 'react-bootstrap';

import axios from 'axios';


function LeagueInformationContainer(props) {
	const [s, setS] = useState(0);

	const playAllWeeks = () => {
		axios.post("http://localhost:8080/league-table").then(
			res => {
				console.log(res.status);
				console.log(res.data);
				
				setS(s+1);
			}
		)
	};

	const playNextWeek = () => {
		axios.post("http://localhost:8080/matches").then(
			res => {
				console.log(res.status);
				console.log(res.data);

				setS(s+1);
			}
		)
	};

	const resetLeague = () => {
		axios.delete("http://localhost:8080/league-table").then(
			res => {
				console.log(res.status);
				console.log(res.data);

				setS(s+1);
			}
		)
	};

	return (
		<div>
			<Row>
				<Col>
					<LeagueTable refresh={s} />
				</Col>
				<Col>
					<MatchesTable refresh={s} />
				</Col>
				<Col>
					<ChampionshipPrediction />
				</Col>
			</Row>
			<Row>
				<Col>
					<Button variant="primary" onClick={playAllWeeks}>Play All Weeks</Button>
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
