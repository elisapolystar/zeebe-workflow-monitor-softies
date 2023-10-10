import React from 'react';

const MessageDisplayer = ({ messages }) => {
    return (
        <div>
            <h1>Messages:</h1>
            <ul>
                {messages.map((message, index) => (
                    <li key={index}>{message}</li>
                ))}
            </ul>
        </div>
  );
};

export default MessageDisplayer;
