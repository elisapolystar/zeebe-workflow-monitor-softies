import React from 'react';

interface ErrorProps {
    errorMessage: string | null; 
  }


const ErrorDisplay: React.FC<ErrorProps> = ({ errorMessage}) => {
    const errorData = errorMessage ? JSON.parse(errorMessage) : {};

    return (
        <div className="error-container">
          <div className="error-item">
            <span>Error</span>
              <div className="error-message" key={errorData.Error}>
                <span>{errorData.Error}</span>
              </div>

            </div> 
        </div> 
      );
  }
  export default ErrorDisplay;