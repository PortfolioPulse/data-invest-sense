import { MinioBucketsExecutorSchema } from './schema';
import { ExecutorContext } from '@nrwl/devkit';
import * as Minio from 'minio';

export default async function runExecutor(
  options: MinioBucketsExecutorSchema,
  context: ExecutorContext,
  ) {
  console.log('Executor ran for MinioBuckets', options);
  const minioClient = new Minio.Client({
    endPoint: 'localhost',
    port: 9000,
    useSSL: false,
    accessKey: 'new-minio-root-user',
    secretKey: 'new-minio-root-password',
  });

  try {
    await minioClient.makeBucket(options.bucketName);
    console.log(`Bucket ${options.bucketName} created successfully`);
    return { success: true };
  } catch (error) {
    console.log(`Error creating bucket: ${error}`);
    return { success: false, error: error.message };
  }
}

