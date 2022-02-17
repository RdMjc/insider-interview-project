import React from 'react';
import LeagueTable from "./LeagueTable";
import MatchesTable from "./MatchesTable";
import ChampionshipPrediction from "./ChampionshipPrediction";
import { Button } from "react-bootstrap";
import { Row, Col } from 'react-bootstrap';


function LeagueInformationContainer(props) {
	return (
		<div>
			<Row>
				<Col>
					<LeagueTable />
				</Col>
				<Col>
					<MatchesTable />
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
					<Button variant="info">Play Next Week</Button>
				</Col>
				<Col>
					<Button variant="danger">Reset</Button>
				</Col>
			</Row>
		</div>
	);
}

export default LeagueInformationContainer;
