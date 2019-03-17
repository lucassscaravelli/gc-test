import React from "react";

const Matches = ({ matches }) => {
  
  return (
    <div>
      <table>
        {matches.map(
          ({
            HostTag,
            HostScore,
            HostColor,
            VisitorTag,
            VisitorScore,
            VisitorColor
          }) => {
            return (
              <tr key={HostTag + VisitorTag}>
                <td
                  style={{
                    background: HostColor,
                    width: "5%"
                  }}
                />
                <td>{HostTag}</td>
                <td>{HostScore}</td>
                <td>X</td>
                <td>{VisitorScore}</td>
                <td>{VisitorTag}</td>
                <td
                  style={{
                    background: VisitorColor,
                    width: "5%"
                  }}
                />
              </tr>
            );
          }
        )}
      </table>
    </div>
  );
};

export default Matches;
