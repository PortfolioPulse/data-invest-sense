import { MinioBucketsExecutorSchema } from './schema';
import executor from './executor';

const options: MinioBucketsExecutorSchema = {};

describe('MinioBuckets Executor', () => {
  it('can run', async () => {
    const output = await executor(options);
    expect(output.success).toBe(true);
  });
});
