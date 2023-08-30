import * as fs from 'fs';
import * as path from 'path';
import axios from 'axios';
import { ExecutorContext } from '@nrwl/devkit';
import { ControllerConfigsExecutorSchema } from './schema';

const LAKE_CONTROLLER_ENDPOINT = 'http://localhost:8002/configs';

export default async function runExecutor(
  options: ControllerConfigsExecutorSchema,
  context: ExecutorContext,
) {
  console.log('Executor ran for ControllerConfigs', options);
  const configsDir = path.join(context.root, 'apps', 'lake-orchestration', 'lake-controller', 'configs');
  try {
    const configFiles = fs.readdirSync(configsDir).filter(file => file.endsWith('.json'));

    for (const configFile of configFiles) {
      const configPath = path.join(configsDir, configFile);
      const jsonBody = JSON.parse(fs.readFileSync(configPath, 'utf-8'));

      try {
        const response = await axios.post(LAKE_CONTROLLER_ENDPOINT, jsonBody);
        console.log(`Posted JSON from ${configFile}:`, response.data);
      } catch (error) {
        console.error(`Error posting JSON from ${configFile}:`, error.message);
      }
    }
  } catch (error) {
    console.error('Error reading configs directory:', error.message);
  }

  return {
    success: true,
  };
}
