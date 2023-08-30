import { ControllerConfigsExecutorSchema } from './schema';
import executor from './executor';

const options: ControllerConfigsExecutorSchema = {};

describe('ControllerConfigs Executor', () => {
  it('can run', async () => {
    const output = await executor(options);
    expect(output.success).toBe(true);
  });
});
