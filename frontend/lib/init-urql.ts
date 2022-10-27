import { initUrqlClient } from 'next-urql';
import { Client } from 'urql';

export function urqlClient(): Promise<Client> {
  return new Promise((resolve, reject) => {
    const client = initUrqlClient(
      {
        url: 'http://localhost:3000/graphql',
      },
      false,
    );
    if (!client) {
      reject(Error('Failed to init initUrqlClient.'));
    } else {
      
      resolve(client);
    }
  });
}
