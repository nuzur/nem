package monitoring

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (i *Implementation) emitLog(req EmitRequest) {
	if i.config.LoggingEnabled {
		extraDataField := parseExtraData(req.ExtraData)
		fields := []zap.Field{
			zap.String("action", req.ActionIdentifier),
			zap.String("entity", req.EntityIdentifier),
			zap.String("layer", string(req.Layer)),
			zap.String("type", string(req.Type)),
		}
		if req.LayerSubtype != "" {
			fields = append(fields, zap.String("layer-subtype", string(req.LayerSubtype)))
		}
		if req.Data != nil {
			om, ok := req.Data.(zapcore.ObjectMarshaler)
			if ok {
				fields = append(fields, zap.Object("data", om))
			} else {
				fields = append(fields, zap.Reflect("data", req.Data))
			}
		}
		fields = append(fields, extraDataField...)
		switch req.Type {
		case EmitTypeError:
			fields = append(fields, zap.Error(req.Error))
			i.logger.Error(
				req.Message,
				fields...,
			)
		case EmitTypeWarn:
			i.logger.Warn(
				req.Message,
				fields...,
			)
		case EmitTypeInfo, EmitTypeSuccess:
			i.logger.Info(
				req.Message,
				fields...,
			)
		}
	}
}

func parseExtraData(data map[string]string) []zap.Field {
	res := []zap.Field{}
	for k, d := range data {
		res = append(res, zap.String(k, d))
	}
	return res
}
