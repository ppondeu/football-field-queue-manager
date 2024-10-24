# StadiumQueueManager

StadiumQueueManager is a robust system designed to manage the scheduling and queueing of teams for sports events, particularly in stadiums with multiple fields. The system allows for efficient queue management, configurable game rules, and team transitions through different states (pending, queue, playing).

## Features

- **Multiple Locations and Stadiums**: Support for multiple stadiums within different locations.
- **Team Queue Management**: Manage teams in a circular queue and handle transitions between pending, queue, and playing states.
- **Configurable Game Logic**:
  - **Unlimited Wins**: Winning team stays in the queue until it loses.
  - **Limited Wins**: Team can play a maximum of two matches before leaving the queue.
  - **Custom Logic**: Easily swap out game logic configurations.
- **Team Registration**: Teams can register and wait in a pending state before being added to the queue.
- **Skip Queue**: Option to skip a team if they are not ready to play.
- **Team State Transitions**: Track teams as they move through different states such as pending, in queue, and playing.
- **Match Completion**: After a match, move teams based on the results (win, lose, draw) following the configured game logic.
