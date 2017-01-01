import React from 'react'

export default class Upload extends React.Component {
  render() {
    return (
      <div className="file-upload">
        <h1>Upload multiple files with fields</h1>

        <form action="http://localhost:8081/upload" method="post">
          Name: <input type="text" name="name" /><br />
          Email: <input type="email" name="email" /><br />
          Files: <input type="file" name="files" /><br /><br />
          <input type="submit" value="Submit" />
        </form>
      </div>
    )
  }
};
