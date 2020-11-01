import 'semantic-ui-css/semantic.min.css'
import React from 'react'
import { render } from 'react-dom'
import { BrowserRouter as Router } from 'react-router-dom'
import {
  Container,
  Grid,
  Segment,
  Header,
  Comment,
  Form,
  Button,
  Image,
  List,
} from 'semantic-ui-react'

if (module.hot) {
  module.hot.dispose(function () {
    // модуль будет заменен.
  })

  module.hot.accept(function () {
    // модуль или одна из его зависимостей была только что обновлена.
  })
}

const App = () => {
  const dialogs = []
  for (let i = 0; i < 20; i++) {
    dialogs.push(
      <List.Item key={`dialog_${i}`}>
        <Image
          avatar
          src='https://react.semantic-ui.com/images/avatar/small/helen.jpg'
        />
        <List.Content>
          <List.Header>Helen</List.Header>
        </List.Content>
      </List.Item>,
    )
  }
  const comments = []
  for (let i = 0; i < 20; i++) {
    comments.push(
      <Comment key={`message_${i}`}>
        <Comment.Avatar
          src='https://react.semantic-ui.com/images/avatar/small/joe.jpg'/>
        <Comment.Content>
          <Comment.Author as='a'>Joe Henderson</Comment.Author>
          <Comment.Metadata>
            <div>5 days ago</div>
          </Comment.Metadata>
          <Comment.Text>Dude, this is awesome. Thanks so
            much</Comment.Text>
          <Comment.Actions>
            <Comment.Action>Reply</Comment.Action>
          </Comment.Actions>
        </Comment.Content>
      </Comment>,
    )
  }

  return <Router>
    <Container style={{ paddingTop: '20px' }}>
      <Grid columns={2} divided>
        <Grid.Row>
          <Grid.Column largeScreen={5}>
            <Segment style={{ height: '95vh' }}>
              <List
                selection
                verticalAlign='middle'
                divided
                style={{
                  overflow: 'auto',
                  maxWidth: '100%',
                  maxHeight: '100%',
                }}
              >
                {dialogs}
              </List>
            </Segment>
          </Grid.Column>
          <Grid.Column largeScreen={11}>
            <Segment basic style={{ height: '90vh' }}>
              <Header as='h3'>
                <Image
                  src='https://react.semantic-ui.com/images/avatar/large/patrick.png'
                />
                Patrick
              </Header>
              <Segment attached='top' style={{ height: '75vh' }}>
                <Comment.Group
                  style={{
                    overflow: 'auto',
                    maxWidth: '100%',
                    maxHeight: '100%',
                  }}
                >
                  {comments}
                </Comment.Group>
              </Segment>
              <Segment attached='bottom'>
                <Form reply>
                  <Form.Input action={
                    <Button
                      icon='caret square right'
                      primary
                    />
                  } placeholder='Write a message...'/>

                </Form>
              </Segment>
            </Segment>
          </Grid.Column>
        </Grid.Row>
      </Grid>
    </Container>
  </Router>
}

render(<App/>, document.getElementById('app'))
