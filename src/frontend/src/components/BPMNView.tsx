import React, { useState, useEffect } from 'react';
import { BpmnVisualization, ShapeBpmnElementKind } from 'bpmn-visualization';
import diagram from './money-loan.bpmn';
import tippy from 'tippy.js';

interface BpmnProps {
    process: string | null;
  }
// const BPMNView: React.FC = () => {
//     const containerRef = useRef<HTMLDivElement>(null);
    
//     return <div className="bpmn-container" ref={containerRef}>BMPN here</div>;
//     }

// export default BPMNView;
const encodedBpmn = "PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz4KPGJwbW46ZGVmaW5pdGlvbnMgeG1sbnM6YnBtbj0iaHR0cDovL3d3dy5vbWcub3JnL3NwZWMvQlBNTi8yMDEwMDUyNC9NT0RFTCIgeG1sbnM6YnBtbmRpPSJodHRwOi8vd3d3Lm9tZy5vcmcvc3BlYy9CUE1OLzIwMTAwNTI0L0RJIiB4bWxuczpkYz0iaHR0cDovL3d3dy5vbWcub3JnL3NwZWMvREQvMjAxMDA1MjQvREMiIHhtbG5zOnhzaT0iaHR0cDovL3d3dy53My5vcmcvMjAwMS9YTUxTY2hlbWEtaW5zdGFuY2UiIHhtbG5zOmRpPSJodHRwOi8vd3d3Lm9tZy5vcmcvc3BlYy9ERC8yMDEwMDUyNC9ESSIgeG1sbnM6emVlYmU9Imh0dHA6Ly9jYW11bmRhLm9yZy9zY2hlbWEvemVlYmUvMS4wIiB4bWxuczptb2RlbGVyPSJodHRwOi8vY2FtdW5kYS5vcmcvc2NoZW1hL21vZGVsZXIvMS4wIiBpZD0iRGVmaW5pdGlvbnNfMGhtdnZqMCIgdGFyZ2V0TmFtZXNwYWNlPSJodHRwOi8vYnBtbi5pby9zY2hlbWEvYnBtbiIgZXhwb3J0ZXI9IkNhbXVuZGEgTW9kZWxlciIgZXhwb3J0ZXJWZXJzaW9uPSI1LjEzLjAiIG1vZGVsZXI6ZXhlY3V0aW9uUGxhdGZvcm09IkNhbXVuZGEgQ2xvdWQiIG1vZGVsZXI6ZXhlY3V0aW9uUGxhdGZvcm1WZXJzaW9uPSI4LjIuMCI+CiAgPGJwbW46cHJvY2VzcyBpZD0ibW9uZXktbG9hbiIgaXNFeGVjdXRhYmxlPSJ0cnVlIj4KICAgIDxicG1uOnN0YXJ0RXZlbnQgaWQ9IlN0YXJ0RXZlbnRfMSIgbmFtZT0iTmVlZGVkIHRvIGxvYW4gbW9uZXkiPgogICAgICA8YnBtbjpvdXRnb2luZz5GbG93XzAzZWkxYjQ8L2JwbW46b3V0Z29pbmc+CiAgICA8L2JwbW46c3RhcnRFdmVudD4KICAgIDxicG1uOmV4Y2x1c2l2ZUdhdGV3YXkgaWQ9IkdhdGV3YXlfMHd2ZW8wYiIgbmFtZT0iQ2hlY2sgY3JlZGl0IiBkZWZhdWx0PSJGbG93XzBrd2s2Y2giPgogICAgICA8YnBtbjppbmNvbWluZz5GbG93XzAzZWkxYjQ8L2JwbW46aW5jb21pbmc+CiAgICAgIDxicG1uOm91dGdvaW5nPkZsb3dfMGt3azZjaDwvYnBtbjpvdXRnb2luZz4KICAgICAgPGJwbW46b3V0Z29pbmc+Rmxvd18xc3czZ3B0PC9icG1uOm91dGdvaW5nPgogICAgPC9icG1uOmV4Y2x1c2l2ZUdhdGV3YXk+CiAgICA8YnBtbjplbmRFdmVudCBpZD0iRXZlbnRfMTlmajN5OSIgbmFtZT0iRW5kIj4KICAgICAgPGJwbW46aW5jb21pbmc+Rmxvd18xZTB6NHZjPC9icG1uOmluY29taW5nPgogICAgICA8YnBtbjppbmNvbWluZz5GbG93XzAybGtkdjQ8L2JwbW46aW5jb21pbmc+CiAgICA8L2JwbW46ZW5kRXZlbnQ+CiAgICA8YnBtbjpzZXJ2aWNlVGFzayBpZD0ic2VuZC1hY2NlcHRhbmNlLWxldHRlciIgbmFtZT0iU2VuZCBhY2NlcHRhbmNlIGxldHRlciI+CiAgICAgIDxicG1uOmV4dGVuc2lvbkVsZW1lbnRzPgogICAgICAgIDx6ZWViZTp0YXNrRGVmaW5pdGlvbiB0eXBlPSJzZW5kLWFjY2VwdGFuY2UtbGV0dGVyIiAvPgogICAgICA8L2JwbW46ZXh0ZW5zaW9uRWxlbWVudHM+CiAgICAgIDxicG1uOmluY29taW5nPkZsb3dfMGt3azZjaDwvYnBtbjppbmNvbWluZz4KICAgICAgPGJwbW46b3V0Z29pbmc+Rmxvd18xYnR3M3JuPC9icG1uOm91dGdvaW5nPgogICAgPC9icG1uOnNlcnZpY2VUYXNrPgogICAgPGJwbW46c2VydmljZVRhc2sgaWQ9InRyYW5zZmVyLW1vbmV5IiBuYW1lPSJUcmFuc2ZlciBtb25leSI+CiAgICAgIDxicG1uOmV4dGVuc2lvbkVsZW1lbnRzPgogICAgICAgIDx6ZWViZTp0YXNrRGVmaW5pdGlvbiB0eXBlPSJ0cmFuc2Zlci1tb25leSIgLz4KICAgICAgPC9icG1uOmV4dGVuc2lvbkVsZW1lbnRzPgogICAgICA8YnBtbjppbmNvbWluZz5GbG93XzFidHczcm48L2JwbW46aW5jb21pbmc+CiAgICAgIDxicG1uOm91dGdvaW5nPkZsb3dfMWUwejR2YzwvYnBtbjpvdXRnb2luZz4KICAgIDwvYnBtbjpzZXJ2aWNlVGFzaz4KICAgIDxicG1uOnNlcnZpY2VUYXNrIGlkPSJzZW5kLXJlamVjdGlvbi1sZXR0ZXIiIG5hbWU9IlNlbmQgcmVqZWN0aW9uIGxldHRlciI+CiAgICAgIDxicG1uOmV4dGVuc2lvbkVsZW1lbnRzPgogICAgICAgIDx6ZWViZTp0YXNrRGVmaW5pdGlvbiB0eXBlPSJzZW5kLXJlamVjdGlvbi1sZXR0ZXIiIC8+CiAgICAgIDwvYnBtbjpleHRlbnNpb25FbGVtZW50cz4KICAgICAgPGJwbW46aW5jb21pbmc+Rmxvd18xc3czZ3B0PC9icG1uOmluY29taW5nPgogICAgICA8YnBtbjppbmNvbWluZz5GbG93XzExNnU3YmI8L2JwbW46aW5jb21pbmc+CiAgICAgIDxicG1uOm91dGdvaW5nPkZsb3dfMXRkMnN3MTwvYnBtbjpvdXRnb2luZz4KICAgIDwvYnBtbjpzZXJ2aWNlVGFzaz4KICAgIDxicG1uOmV4Y2x1c2l2ZUdhdGV3YXkgaWQ9IkdhdGV3YXlfMDdqbGM5ayIgbmFtZT0iQ3VzdG9tZXIgcmVjZWl2ZSBsZXR0ZXI/IiBkZWZhdWx0PSJGbG93XzFtNjE4OTgiPgogICAgICA8YnBtbjppbmNvbWluZz5GbG93XzF0ZDJzdzE8L2JwbW46aW5jb21pbmc+CiAgICAgIDxicG1uOm91dGdvaW5nPkZsb3dfMDJsa2R2NDwvYnBtbjpvdXRnb2luZz4KICAgICAgPGJwbW46b3V0Z29pbmc+Rmxvd18xbTYxODk4PC9icG1uOm91dGdvaW5nPgogICAgPC9icG1uOmV4Y2x1c2l2ZUdhdGV3YXk+CiAgICA8YnBtbjppbnRlcm1lZGlhdGVDYXRjaEV2ZW50IGlkPSJFdmVudF8wOXZldzRtIiBuYW1lPSJXYWl0IDIgbWludXRlcyI+CiAgICAgIDxicG1uOmluY29taW5nPkZsb3dfMW02MTg5ODwvYnBtbjppbmNvbWluZz4KICAgICAgPGJwbW46b3V0Z29pbmc+Rmxvd18xMTZ1N2JiPC9icG1uOm91dGdvaW5nPgogICAgICA8YnBtbjp0aW1lckV2ZW50RGVmaW5pdGlvbiBpZD0iVGltZXJFdmVudERlZmluaXRpb25fMWsyeTNjZCI+CiAgICAgICAgPGJwbW46dGltZUR1cmF0aW9uIHhzaTp0eXBlPSJicG1uOnRGb3JtYWxFeHByZXNzaW9uIj5QVDEyMFM8L2JwbW46dGltZUR1cmF0aW9uPgogICAgICA8L2JwbW46dGltZXJFdmVudERlZmluaXRpb24+CiAgICA8L2JwbW46aW50ZXJtZWRpYXRlQ2F0Y2hFdmVudD4KICAgIDxicG1uOnNlcXVlbmNlRmxvdyBpZD0iRmxvd18wM2VpMWI0IiBzb3VyY2VSZWY9IlN0YXJ0RXZlbnRfMSIgdGFyZ2V0UmVmPSJHYXRld2F5XzB3dmVvMGIiIC8+CiAgICA8YnBtbjpzZXF1ZW5jZUZsb3cgaWQ9IkZsb3dfMGt3azZjaCIgc291cmNlUmVmPSJHYXRld2F5XzB3dmVvMGIiIHRhcmdldFJlZj0ic2VuZC1hY2NlcHRhbmNlLWxldHRlciIgLz4KICAgIDxicG1uOnNlcXVlbmNlRmxvdyBpZD0iRmxvd18xc3czZ3B0IiBzb3VyY2VSZWY9IkdhdGV3YXlfMHd2ZW8wYiIgdGFyZ2V0UmVmPSJzZW5kLXJlamVjdGlvbi1sZXR0ZXIiPgogICAgICA8YnBtbjpjb25kaXRpb25FeHByZXNzaW9uIHhzaTp0eXBlPSJicG1uOnRGb3JtYWxFeHByZXNzaW9uIj49ZGVidCAmZ3Q7IDEwMDA8L2JwbW46Y29uZGl0aW9uRXhwcmVzc2lvbj4KICAgIDwvYnBtbjpzZXF1ZW5jZUZsb3c+CiAgICA8YnBtbjpzZXF1ZW5jZUZsb3cgaWQ9IkZsb3dfMWUwejR2YyIgc291cmNlUmVmPSJ0cmFuc2Zlci1tb25leSIgdGFyZ2V0UmVmPSJFdmVudF8xOWZqM3k5IiAvPgogICAgPGJwbW46c2VxdWVuY2VGbG93IGlkPSJGbG93XzAybGtkdjQiIHNvdXJjZVJlZj0iR2F0ZXdheV8wN2psYzlrIiB0YXJnZXRSZWY9IkV2ZW50XzE5ZmozeTkiPgogICAgICA8YnBtbjpjb25kaXRpb25FeHByZXNzaW9uIHhzaTp0eXBlPSJicG1uOnRGb3JtYWxFeHByZXNzaW9uIj49cmVjZWl2ZWQgPSB0cnVlPC9icG1uOmNvbmRpdGlvbkV4cHJlc3Npb24+CiAgICA8L2JwbW46c2VxdWVuY2VGbG93PgogICAgPGJwbW46c2VxdWVuY2VGbG93IGlkPSJGbG93XzFidHczcm4iIHNvdXJjZVJlZj0ic2VuZC1hY2NlcHRhbmNlLWxldHRlciIgdGFyZ2V0UmVmPSJ0cmFuc2Zlci1tb25leSIgLz4KICAgIDxicG1uOnNlcXVlbmNlRmxvdyBpZD0iRmxvd18xMTZ1N2JiIiBzb3VyY2VSZWY9IkV2ZW50XzA5dmV3NG0iIHRhcmdldFJlZj0ic2VuZC1yZWplY3Rpb24tbGV0dGVyIiAvPgogICAgPGJwbW46c2VxdWVuY2VGbG93IGlkPSJGbG93XzFtNjE4OTgiIHNvdXJjZVJlZj0iR2F0ZXdheV8wN2psYzlrIiB0YXJnZXRSZWY9IkV2ZW50XzA5dmV3NG0iIC8+CiAgICA8YnBtbjpzZXF1ZW5jZUZsb3cgaWQ9IkZsb3dfMXRkMnN3MSIgc291cmNlUmVmPSJzZW5kLXJlamVjdGlvbi1sZXR0ZXIiIHRhcmdldFJlZj0iR2F0ZXdheV8wN2psYzlrIiAvPgogIDwvYnBtbjpwcm9jZXNzPgogIDxicG1uZGk6QlBNTkRpYWdyYW0gaWQ9IkJQTU5EaWFncmFtXzEiPgogICAgPGJwbW5kaTpCUE1OUGxhbmUgaWQ9IkJQTU5QbGFuZV8xIiBicG1uRWxlbWVudD0ibW9uZXktbG9hbiI+CiAgICAgIDxicG1uZGk6QlBNTlNoYXBlIGlkPSJCUE1OU2hhcGVfMHNrbHlyciIgYnBtbkVsZW1lbnQ9IlN0YXJ0RXZlbnRfMSI+CiAgICAgICAgPGRjOkJvdW5kcyB4PSIxNzIiIHk9IjE4MiIgd2lkdGg9IjM2IiBoZWlnaHQ9IjM2IiAvPgogICAgICAgIDxicG1uZGk6QlBNTkxhYmVsPgogICAgICAgICAgPGRjOkJvdW5kcyB4PSIxNTMiIHk9IjIyNSIgd2lkdGg9Ijc1IiBoZWlnaHQ9IjI3IiAvPgogICAgICAgIDwvYnBtbmRpOkJQTU5MYWJlbD4KICAgICAgPC9icG1uZGk6QlBNTlNoYXBlPgogICAgICA8YnBtbmRpOkJQTU5TaGFwZSBpZD0iQlBNTlNoYXBlXzAydmthbXgiIGJwbW5FbGVtZW50PSJHYXRld2F5XzB3dmVvMGIiIGlzTWFya2VyVmlzaWJsZT0idHJ1ZSI+CiAgICAgICAgPGRjOkJvdW5kcyB4PSIyOTUiIHk9IjE3NSIgd2lkdGg9IjUwIiBoZWlnaHQ9IjUwIiAvPgogICAgICAgIDxicG1uZGk6QlBNTkxhYmVsPgogICAgICAgICAgPGRjOkJvdW5kcyB4PSIzNTUiIHk9IjE5MyIgd2lkdGg9IjYyIiBoZWlnaHQ9IjE0IiAvPgogICAgICAgIDwvYnBtbmRpOkJQTU5MYWJlbD4KICAgICAgPC9icG1uZGk6QlBNTlNoYXBlPgogICAgICA8YnBtbmRpOkJQTU5TaGFwZSBpZD0iQlBNTlNoYXBlXzBtYTIwbzAiIGJwbW5FbGVtZW50PSJFdmVudF8xOWZqM3k5Ij4KICAgICAgICA8ZGM6Qm91bmRzIHg9Ijg5MiIgeT0iMTgyIiB3aWR0aD0iMzYiIGhlaWdodD0iMzYiIC8+CiAgICAgICAgPGJwbW5kaTpCUE1OTGFiZWw+CiAgICAgICAgICA8ZGM6Qm91bmRzIHg9Ijg2MiIgeT0iMTkzIiB3aWR0aD0iMjAiIGhlaWdodD0iMTQiIC8+CiAgICAgICAgPC9icG1uZGk6QlBNTkxhYmVsPgogICAgICA8L2JwbW5kaTpCUE1OU2hhcGU+CiAgICAgIDxicG1uZGk6QlBNTlNoYXBlIGlkPSJCUE1OU2hhcGVfMHZtZ3l5diIgYnBtbkVsZW1lbnQ9InNlbmQtYWNjZXB0YW5jZS1sZXR0ZXIiPgogICAgICAgIDxkYzpCb3VuZHMgeD0iNDUwIiB5PSI4MCIgd2lkdGg9IjEwMCIgaGVpZ2h0PSI4MCIgLz4KICAgICAgPC9icG1uZGk6QlBNTlNoYXBlPgogICAgICA8YnBtbmRpOkJQTU5TaGFwZSBpZD0iQlBNTlNoYXBlXzFhajIyYTciIGJwbW5FbGVtZW50PSJ0cmFuc2Zlci1tb25leSI+CiAgICAgICAgPGRjOkJvdW5kcyB4PSI2NjAiIHk9IjgwIiB3aWR0aD0iMTAwIiBoZWlnaHQ9IjgwIiAvPgogICAgICA8L2JwbW5kaTpCUE1OU2hhcGU+CiAgICAgIDxicG1uZGk6QlBNTlNoYXBlIGlkPSJCUE1OU2hhcGVfMTBndzM0NyIgYnBtbkVsZW1lbnQ9InNlbmQtcmVqZWN0aW9uLWxldHRlciI+CiAgICAgICAgPGRjOkJvdW5kcyB4PSI0NDAiIHk9IjI1MCIgd2lkdGg9IjEwMCIgaGVpZ2h0PSI4MCIgLz4KICAgICAgPC9icG1uZGk6QlBNTlNoYXBlPgogICAgICA8YnBtbmRpOkJQTU5TaGFwZSBpZD0iQlBNTlNoYXBlXzBrMW42aGYiIGJwbW5FbGVtZW50PSJFdmVudF8wOXZldzRtIj4KICAgICAgICA8ZGM6Qm91bmRzIHg9IjYxMiIgeT0iMzgyIiB3aWR0aD0iMzYiIGhlaWdodD0iMzYiIC8+CiAgICAgICAgPGJwbW5kaTpCUE1OTGFiZWw+CiAgICAgICAgICA8ZGM6Qm91bmRzIHg9IjU5NCIgeT0iNDI1IiB3aWR0aD0iNzMiIGhlaWdodD0iMTQiIC8+CiAgICAgICAgPC9icG1uZGk6QlBNTkxhYmVsPgogICAgICA8L2JwbW5kaTpCUE1OU2hhcGU+CiAgICAgIDxicG1uZGk6QlBNTlNoYXBlIGlkPSJCUE1OU2hhcGVfMGRzbWxjMSIgYnBtbkVsZW1lbnQ9IkdhdGV3YXlfMDdqbGM5ayIgaXNNYXJrZXJWaXNpYmxlPSJ0cnVlIj4KICAgICAgICA8ZGM6Qm91bmRzIHg9Ijc1NSIgeT0iMjY1IiB3aWR0aD0iNTAiIGhlaWdodD0iNTAiIC8+CiAgICAgICAgPGJwbW5kaTpCUE1OTGFiZWw+CiAgICAgICAgICA8ZGM6Qm91bmRzIHg9IjczNyIgeT0iMjM1IiB3aWR0aD0iODciIGhlaWdodD0iMjciIC8+CiAgICAgICAgPC9icG1uZGk6QlBNTkxhYmVsPgogICAgICA8L2JwbW5kaTpCUE1OU2hhcGU+CiAgICAgIDxicG1uZGk6QlBNTkVkZ2UgaWQ9IkJQTU5FZGdlXzBsbzNxY2EiIGJwbW5FbGVtZW50PSJGbG93XzAzZWkxYjQiPgogICAgICAgIDxkaTp3YXlwb2ludCB4PSIyMDgiIHk9IjIwMCIgLz4KICAgICAgICA8ZGk6d2F5cG9pbnQgeD0iMjk1IiB5PSIyMDAiIC8+CiAgICAgIDwvYnBtbmRpOkJQTU5FZGdlPgogICAgICA8YnBtbmRpOkJQTU5FZGdlIGlkPSJCUE1ORWRnZV8wZDBpemE0IiBicG1uRWxlbWVudD0iRmxvd18wa3drNmNoIj4KICAgICAgICA8ZGk6d2F5cG9pbnQgeD0iMzIwIiB5PSIxNzUiIC8+CiAgICAgICAgPGRpOndheXBvaW50IHg9IjMyMCIgeT0iMTIwIiAvPgogICAgICAgIDxkaTp3YXlwb2ludCB4PSI0NTAiIHk9IjEyMCIgLz4KICAgICAgPC9icG1uZGk6QlBNTkVkZ2U+CiAgICAgIDxicG1uZGk6QlBNTkVkZ2UgaWQ9IkJQTU5FZGdlXzF2YTN3MDgiIGJwbW5FbGVtZW50PSJGbG93XzFzdzNncHQiPgogICAgICAgIDxkaTp3YXlwb2ludCB4PSIzMjAiIHk9IjIyNSIgLz4KICAgICAgICA8ZGk6d2F5cG9pbnQgeD0iMzIwIiB5PSIyOTAiIC8+CiAgICAgICAgPGRpOndheXBvaW50IHg9IjQ0MCIgeT0iMjkwIiAvPgogICAgICA8L2JwbW5kaTpCUE1ORWRnZT4KICAgICAgPGJwbW5kaTpCUE1ORWRnZSBpZD0iQlBNTkVkZ2VfMGEycjNzOSIgYnBtbkVsZW1lbnQ9IkZsb3dfMWUwejR2YyI+CiAgICAgICAgPGRpOndheXBvaW50IHg9Ijc2MCIgeT0iMTIwIiAvPgogICAgICAgIDxkaTp3YXlwb2ludCB4PSI5MTAiIHk9IjEyMCIgLz4KICAgICAgICA8ZGk6d2F5cG9pbnQgeD0iOTEwIiB5PSIxODIiIC8+CiAgICAgIDwvYnBtbmRpOkJQTU5FZGdlPgogICAgICA8YnBtbmRpOkJQTU5FZGdlIGlkPSJCUE1ORWRnZV8xeW93MHAyIiBicG1uRWxlbWVudD0iRmxvd18wMmxrZHY0Ij4KICAgICAgICA8ZGk6d2F5cG9pbnQgeD0iODA1IiB5PSIyOTAiIC8+CiAgICAgICAgPGRpOndheXBvaW50IHg9IjkxMCIgeT0iMjkwIiAvPgogICAgICAgIDxkaTp3YXlwb2ludCB4PSI5MTAiIHk9IjIxOCIgLz4KICAgICAgPC9icG1uZGk6QlBNTkVkZ2U+CiAgICAgIDxicG1uZGk6QlBNTkVkZ2UgaWQ9IkJQTU5FZGdlXzFyNXh1ODEiIGJwbW5FbGVtZW50PSJGbG93XzFidHczcm4iPgogICAgICAgIDxkaTp3YXlwb2ludCB4PSI1NTAiIHk9IjEyMCIgLz4KICAgICAgICA8ZGk6d2F5cG9pbnQgeD0iNjYwIiB5PSIxMjAiIC8+CiAgICAgIDwvYnBtbmRpOkJQTU5FZGdlPgogICAgICA8YnBtbmRpOkJQTU5FZGdlIGlkPSJCUE1ORWRnZV8waHM2bmh5IiBicG1uRWxlbWVudD0iRmxvd18xMTZ1N2JiIj4KICAgICAgICA8ZGk6d2F5cG9pbnQgeD0iNjEyIiB5PSI0MDAiIC8+CiAgICAgICAgPGRpOndheXBvaW50IHg9IjQ5MCIgeT0iNDAwIiAvPgogICAgICAgIDxkaTp3YXlwb2ludCB4PSI0OTAiIHk9IjMzMCIgLz4KICAgICAgPC9icG1uZGk6QlBNTkVkZ2U+CiAgICAgIDxicG1uZGk6QlBNTkVkZ2UgaWQ9IkJQTU5FZGdlXzE2MWJnamQiIGJwbW5FbGVtZW50PSJGbG93XzFtNjE4OTgiPgogICAgICAgIDxkaTp3YXlwb2ludCB4PSI3ODAiIHk9IjMxNSIgLz4KICAgICAgICA8ZGk6d2F5cG9pbnQgeD0iNzgwIiB5PSI0MDAiIC8+CiAgICAgICAgPGRpOndheXBvaW50IHg9IjY0OCIgeT0iNDAwIiAvPgogICAgICA8L2JwbW5kaTpCUE1ORWRnZT4KICAgICAgPGJwbW5kaTpCUE1ORWRnZSBpZD0iRmxvd18xdGQyc3cxX2RpIiBicG1uRWxlbWVudD0iRmxvd18xdGQyc3cxIj4KICAgICAgICA8ZGk6d2F5cG9pbnQgeD0iNTQwIiB5PSIyOTAiIC8+CiAgICAgICAgPGRpOndheXBvaW50IHg9Ijc1NSIgeT0iMjkwIiAvPgogICAgICA8L2JwbW5kaTpCUE1ORWRnZT4KICAgIDwvYnBtbmRpOkJQTU5QbGFuZT4KICA8L2JwbW5kaTpCUE1ORGlhZ3JhbT4KPC9icG1uOmRlZmluaXRpb25zPgo=";
const decodedBpmn = atob(encodedBpmn);
console.log(decodedBpmn);

const BPMNView: React.FC<BpmnProps> = () => {
  const [diagramData, setDiagramData] = useState<string | null>(null);
  const bpmnContainerElt = window.document.getElementById('bpmn-container');

  useEffect(() => {
      async function fetchData() {
          try {
              const response = await fetch(diagram);
              const data = await response.text();
              setDiagramData(data);
          } catch (error) {
              console.error('Error fetching diagram:', error);
          }
      }

      fetchData();
  }, []);

  useEffect(() => {
      if (diagramData) {
          const bpmnVisualization = new BpmnVisualization({ container: bpmnContainerElt as HTMLElement, navigation: { enabled: false } });
          bpmnVisualization.load(diagramData);
          const bpmnElementsRegistry = bpmnVisualization.bpmnElementsRegistry;
          const elementsWithTimerPopup = bpmnElementsRegistry.getElementsByKinds(ShapeBpmnElementKind.EVENT_INTERMEDIATE_CATCH);
          const elementsWithPopup = bpmnElementsRegistry.getElementsByKinds(ShapeBpmnElementKind.TASK_SERVICE);
          addPopup(elementsWithPopup);
          console.log(elementsWithTimerPopup);
          addTimerPopup(elementsWithTimerPopup);
      }

      tippy.setDefaultProps({
          content: 'Loading...',
          allowHTML: true,
          onShow(instance) {
              instance.setContent(getBpmnElementInfoAsHtml(instance.reference));
          },
          onHidden(instance) {
              instance.setContent('Loading...');
          },

          // don't consider `data-tippy-*` attributes on the reference element as we fully manage tippy with javascript
          // and we cannot update the reference here as it is generated by bpmn-visualization
          ignoreAttributes: true,

          // https://atomiks.github.io/tippyjs/v6/all-props/#interactive
          interactive: true,
      });

      function addPopup(bpmnElements) {
          bpmnElements.forEach(bpmnElement => {
              const htmlElement = bpmnElement.htmlElement;

              htmlElement.addEventListener('click', () => {
                console.log('Element clicked:', bpmnElement);
            });

              const isEdge = !bpmnElement.bpmnSemantic.isShape;
              const offset = isEdge ? [0, -40] : undefined; // undefined offset for tippyjs default offset

              if (bpmnContainerElt) {
                tippy(htmlElement, {
                  // work perfectly on hover with or without 'diagram navigation' enable
                  appendTo: bpmnContainerElt.parentElement as Element,
                  arrow: false,
                  offset: offset as [number, number] | undefined,
                  placement: 'bottom',
                });
              }
          });
      }
      
      function addTimerPopup(bpmnElements) {
        bpmnElements.forEach(bpmnElement => {
            console.log(bpmnElement);
        });
    }

      function getBpmnElementInfoAsHtml(htmlElement) {
          const rect = htmlElement.querySelector('rect');
          const height = rect.getAttribute('height');
          const width = rect.getAttribute('width');

          return `<div class="bpmn-popover" style="background-color: blue; color: white; padding: 10px; border-radius: 5px; opacity: 0.85;">
      BPMN Info
      <hr>
      box heigth: ${height}
      <br>
      box width: ${width}
      </div>`;
      }

  }, [bpmnContainerElt, diagramData]);

  return (
      <div id="bpmn-container">
          <h2>BPMN demo</h2>
      </div>
  );
}

export default BPMNView;

/*default="Flow_1m61898">
      <bpmn:incoming>Flow_1td2sw1</bpmn:incoming>
      <bpmn:outgoing>Flow_02lkdv4</bpmn:outgoing>
      <bpmn:outgoing>Flow_1m61898</bpmn:outgoing>
    </bpmn:exclusiveGateway>
    <bpmn:intermediateCatchEvent id="Event_09vew4m" name="Wait 2 minutes">
      <bpmn:incoming>Flow_1m61898</bpmn:incoming>
      <bpmn:outgoing>Flow_116u7bb</bpmn:outgoing>
      <bpmn:timerEventDefinition id="TimerEventDefinition_1k2y3cd">
        <bpmn:timeDuration xsi:type="bpmn:tFormalExpression">PT120S</bpmn:timeDuration>
      </bpmn:timerEventDefinition> */